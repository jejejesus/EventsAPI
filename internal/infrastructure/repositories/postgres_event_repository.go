package repositories

import (
	"EventsAPI/internal/domain/entities"
	"EventsAPI/internal/domain/repositories"
	"context"

	"gorm.io/gorm"
)

type postgresEventRepository struct {
	db *gorm.DB
}

func NewPostgresEventRepository(db *gorm.DB) repositories.EventRepository {
	return &postgresEventRepository{db: db}
}

func (r *postgresEventRepository) Create(ctx context.Context, event *entities.Event) error {
	return r.db.WithContext(ctx).Create(event).Error
}

func (r *postgresEventRepository) GetByID(ctx context.Context, id uint) (*entities.Event, error) {
	var event entities.Event
	err := r.db.WithContext(ctx).Preload("User").Preload("Attendees.User").First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *postgresEventRepository) GetByUserID(ctx context.Context, userID uint, limit, offset int) ([]*entities.Event, error) {
	var events []*entities.Event
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Preload("User").
		Limit(limit).
		Offset(offset).
		Find(&events).Error
	return events, err
}

func (r *postgresEventRepository) Update(ctx context.Context, event *entities.Event) error {
	return r.db.WithContext(ctx).Save(event).Error
}

func (r *postgresEventRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Event{}, id).Error
}

func (r *postgresEventRepository) List(ctx context.Context, limit, offset int) ([]*entities.Event, error) {
	var events []*entities.Event
	err := r.db.WithContext(ctx).
		Preload("User").
		Limit(limit).
		Offset(offset).
		Find(&events).Error
	return events, err
}
