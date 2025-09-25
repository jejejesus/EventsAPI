package database

import (
	"fmt"
	"log"

	"EventsAPI/internal/config"
	"EventsAPI/internal/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	// Auto-migrate tables
	err = db.AutoMigrate(
		&entities.User{},
		&entities.Event{},
		&entities.Attendee{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	log.Println("✅ Database connected and migrated successfully")
	return db, nil
}
