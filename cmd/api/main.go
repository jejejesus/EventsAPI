package main

import (
	"log"

	"EventsAPI/docs"
	"EventsAPI/internal/config"
	"EventsAPI/internal/delivery/http/handlers"
	"EventsAPI/internal/delivery/http/routes"
	"EventsAPI/internal/infrastructure/database"
	"EventsAPI/internal/infrastructure/repositories"
	"EventsAPI/internal/usecases"

	"gorm.io/gorm"
)

// @title Events API
// @version 1.0
// @description A complete REST API for event management with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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
	// Initialize repositories
	userRepo := repositories.NewPostgresUserRepository(db)
	eventRepo := repositories.NewPostgresEventRepository(db)
	attendeeRepo := repositories.NewPostgresAttendeeRepository(db)

	// Initialize use cases
	authUseCase := usecases.NewAuthUseCase(userRepo, configs)
	eventUseCase := usecases.NewEventUseCase(eventRepo, userRepo)
	attendeeUseCase := usecases.NewAttendeeUseCase(attendeeRepo, eventRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authUseCase)
	eventHandler := handlers.NewEventHandler(eventUseCase)
	attendeeHandler := handlers.NewAttendeeHandler(attendeeUseCase)
	healthHandler := handlers.NewHealthHandler()

	// Setup routes
	router := routes.SetupRoutes(configs, authHandler, eventHandler, attendeeHandler, healthHandler)

	// Start server
	log.Printf("🚀 Server starting on port %s", configs.Server.Port)
	log.Printf("📚 Swagger documentation available at: http://localhost:%s/swagger/index.html", configs.Server.Port)

	if err := router.Run(":" + configs.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
