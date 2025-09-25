package entities

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	Location    string         `json:"location" gorm:"not null"`
	DateTime    time.Time      `json:"date_time" gorm:"not null"`
	MaxCapacity int            `json:"max_capacity" gorm:"default:0"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	User        User           `json:"user" gorm:"foreignKey:UserID"`
	Attendees   []Attendee     `json:"attendees" gorm:"foreignKey:EventID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
type EventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	MaxCapacity int       `json:"max_capacity" binding:"min=0"`
}
type EventResponse struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Location       string    `json:"location"`
	DateTime       time.Time `json:"date_time"`
	MaxCapacity    int       `json:"max_capacity"`
	UserID         uint      `json:"user_id"`
	AttendeesCount int       `json:"attendees_count"`
	CreatedAt      time.Time `json:"created_at"`
}
