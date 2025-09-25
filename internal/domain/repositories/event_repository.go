package repositories

import (
	"EventsAPI/internal/domain/entities"
	"context"
)

type EventRepository interface {
	Create(ctx context.Context, event *entities.Event) error
	GetByID(ctx context.Context, id uint) (*entities.Event, error)
	GetByUserID(ctx context.Context, userID uint, limit, offset int) ([]*entities.Event, error)
	Update(ctx context.Context, event *entities.Event) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*entities.Event, error)
}
