package usecases

import (
	"EventsAPI/internal/domain/entities"
	"EventsAPI/internal/domain/repositories"
	"context"
	"errors"
)

type AttendeeUseCase struct {
	attendeeRepo repositories.AttendeeRepository
	eventRepo    repositories.EventRepository
}

func NewAttendeeUseCase(attendeeRepo repositories.AttendeeRepository, eventRepo repositories.EventRepository) *AttendeeUseCase {
	return &AttendeeUseCase{attendeeRepo: attendeeRepo, eventRepo: eventRepo}
}

func (uc *AttendeeUseCase) RegisterForEvent(ctx context.Context, eventID, userID uint) error {
	event, err := uc.eventRepo.GetByID(ctx, eventID)
	if err != nil {
		return errors.New("evento no existe")
	}

	count, err := uc.attendeeRepo.CountByEventID(ctx, eventID)
	if err != nil {
		return err
	}

	if int(count) >= event.MaxCapacity {
		return errors.New("no hay cupo disponible en el evento")
	}

	exists, err := uc.attendeeRepo.IsUserRegistered(ctx, eventID, userID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("el usuario ya está registrado")
	}

	attendee := &entities.Attendee{EventID: eventID, UserID: userID}
	return uc.attendeeRepo.Create(ctx, attendee)
}

func (uc *AttendeeUseCase) UnregisterFromEvent(ctx context.Context, eventID, userID uint) error {
	return uc.attendeeRepo.Delete(ctx, eventID, userID)
}

func (uc *AttendeeUseCase) GetMyRegistrations(ctx context.Context, userID uint, limit, offset int) ([]*entities.Attendee, error) {
	return uc.attendeeRepo.GetByUserID(ctx, userID, limit, offset)
}

func (uc *AttendeeUseCase) GetEventAttendees(ctx context.Context, eventID uint, limit, offset int) ([]*entities.Attendee, error) {
	return uc.attendeeRepo.GetByEventID(ctx, eventID, limit, offset)
}
