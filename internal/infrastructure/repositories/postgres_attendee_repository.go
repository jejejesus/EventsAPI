package repositories

import (
	"EventsAPI/internal/domain/entities"
	"EventsAPI/internal/domain/repositories"
	"context"

	"gorm.io/gorm"
)

type postgresAttendeeRepository struct {
	db *gorm.DB
}

func NewPostgresAttendeeRepository(db *gorm.DB) repositories.AttendeeRepository {
	return &postgresAttendeeRepository{db: db}
}

func (r *postgresAttendeeRepository) Create(ctx context.Context, attendee *entities.Attendee) error {
	return r.db.WithContext(ctx).Create(attendee).Error
}

func (r *postgresAttendeeRepository) GetByID(ctx context.Context, id uint) (*entities.Attendee, error) {
	var attendee entities.Attendee
	err := r.db.WithContext(ctx).First(&attendee, id).Error
	if err != nil {
		return nil, err
	}
	return &attendee, nil
}

func (r *postgresAttendeeRepository) GetByEventID(ctx context.Context, eventID uint, limit, offset int) ([]*entities.Attendee, error) {
	var attendees []*entities.Attendee
	err := r.db.WithContext(ctx).Where("event_id = ?", eventID).Limit(limit).Offset(offset).Find(&attendees).Error
	if err != nil {
		return nil, err
	}
	return attendees, nil
}

func (r *postgresAttendeeRepository) GetByUserID(ctx context.Context, userID uint, limit, offset int) ([]*entities.Attendee, error) {
	var attendees []*entities.Attendee
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&attendees).Error
	if err != nil {
		return nil, err
	}
	return attendees, nil
}

func (r *postgresAttendeeRepository) Delete(ctx context.Context, eventID, userID uint) error {
	return r.db.WithContext(ctx).Where("event_id = ? AND user_id = ?", eventID, userID).Delete(&entities.Attendee{}).Error
}

func (r *postgresAttendeeRepository) IsUserRegistered(ctx context.Context, eventID, userID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entities.Attendee{}).Where("event_id = ? AND user_id = ?", eventID, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *postgresAttendeeRepository) CountByEventID(ctx context.Context, eventID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entities.Attendee{}).Where("event_id = ?", eventID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
