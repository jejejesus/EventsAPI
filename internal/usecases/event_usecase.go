package usecases

import (
	"EventsAPI/internal/domain/entities"
	"EventsAPI/internal/domain/repositories"
	"context"
	"errors"
)

type EventUseCase struct {
	eventRepo repositories.EventRepository
	userRepo  repositories.UserRepository
}

func NewEventUseCase(eventRepo repositories.EventRepository, userRepo repositories.UserRepository) *EventUseCase {
	return &EventUseCase{eventRepo: eventRepo, userRepo: userRepo}
}

func (uc *EventUseCase) CreateEvent(ctx context.Context, event *entities.Event) error {
	user, err := uc.userRepo.GetByID(ctx, event.UserID)
	if err != nil {
		return errors.New("usuario no encontrado")
	}

	if event.MaxCapacity < 1 {
		return errors.New("la capacidad del evento debe ser mayor que cero")
	}

	// Validaciones adicionales...
	event.User = *user
	return uc.eventRepo.Create(ctx, event)
}

func (uc *EventUseCase) ListEvents(ctx context.Context, limit, offset int) ([]*entities.Event, error) {
	return uc.eventRepo.List(ctx, limit, offset)
}

func (uc *EventUseCase) GetEventByID(ctx context.Context, id uint) (*entities.Event, error) {
	return uc.eventRepo.GetByID(ctx, id)
}

func (uc *EventUseCase) UpdateEvent(ctx context.Context, event *entities.Event) error {
	return uc.eventRepo.Update(ctx, event)
}

func (uc *EventUseCase) DeleteEvent(ctx context.Context, id uint) error {
	return uc.eventRepo.Delete(ctx, id)
}

func (uc *EventUseCase) GetUserEvents(ctx context.Context, userID uint, limit, offset int) ([]*entities.Event, error) {
	return uc.eventRepo.GetByUserID(ctx, userID, limit, offset)
}
