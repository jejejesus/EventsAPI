package repositories

import (
	"EventsAPI/internal/domain/entities"
	"context"
)

type AttendeeRepository interface {
	Create(ctx context.Context, attendee *entities.Attendee) error
	GetByID(ctx context.Context, id uint) (*entities.Attendee, error)
	GetByEventID(ctx context.Context, eventID uint, limit, offset int) ([]*entities.Attendee, error)
	GetByUserID(ctx context.Context, userID uint, limit, offset int) ([]*entities.Attendee, error)
	Delete(ctx context.Context, eventID, userID uint) error
	IsUserRegistered(ctx context.Context, eventID, userID uint) (bool, error)
	CountByEventID(ctx context.Context, eventID uint) (int64, error)
}
