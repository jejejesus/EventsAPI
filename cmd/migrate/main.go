package main

import (
	"fmt"
	"log"

	"EventsAPI/docs"
	"EventsAPI/internal/config"
	"EventsAPI/internal/domain/entities"
	"EventsAPI/internal/infrastructure/database"

	"gorm.io/gorm"
)

var configs *config.Config
var db *gorm.DB

func init() {
	// Load configuration
	var err error
	configs, err = config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Set Swagger info
	docs.SwaggerInfo.Host = "localhost:" + configs.Server.Port
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Initialize database
	db, err = database.NewPostgresConnection(configs)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func main() {
	fmt.Println("Running database migrations...")

	err := db.AutoMigrate(
		&entities.User{},
		&entities.Event{},
		&entities.Attendee{},
	)
	if err != nil {
		log.Fatalf("could not migrate database: %v", err)
	}

	fmt.Println("Database migration completed successfully!")
}
