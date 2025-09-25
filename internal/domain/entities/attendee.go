package entities

import (
	"time"

	"gorm.io/gorm"
)

type Attendee struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	EventID   uint           `json:"event_id" gorm:"not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	Event     Event          `json:"event" gorm:"foreignKey:EventID"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Constraint: unique combination of EventID and UserID
// This should be added in database migration or GORM constraint
type AttendeeRequest struct {
	EventID uint `json:"event_id" binding:"required"`
}
type AttendeeResponse struct {
	ID      uint          `json:"id"`
	EventID uint          `json:"event_id"`
	UserID  uint          `json:"user_id"`
	Event   EventResponse `json:"event,omitempty"`
	User    UserResponse  `json:"user,omitempty"`
}
