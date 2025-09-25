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

func main() {
	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Set Swagger info
	docs.SwaggerInfo.Host = "localhost:" + config.Server.Port
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Initialize database
	db, err := database.NewPostgresConnection(config)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repositories
	userRepo := repositories.NewPostgresUserRepository(db)
	/*eventRepo := repositories.NewPostgresEventRepository(db)
	attendeeRepo := repositories.NewPostgresAttendeeRepository(db)*/

	// Initialize use cases
	authUseCase := usecases.NewAuthUseCase(userRepo, config)
	/*eventUseCase := usecases.NewEventUseCase(eventRepo, userRepo)
	attendeeUseCase := usecases.NewAttendeeUseCase(attendeeRepo, eventRepo)*/

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authUseCase)
	/*eventHandler := handlers.NewEventHandler(eventUseCase)
	attendeeHandler := handlers.NewAttendeeHandler(attendeeUseCase)*/

	// Setup routes
	router := routes.SetupRoutes(config, authHandler) //, eventHandler, attendeeHandler)

	// Start server
	log.Printf("🚀 Server starting on port %s", config.Server.Port)
	log.Printf("📚 Swagger documentation available at: http://localhost:%s/swagger/index.html", config.Server.Port)

	if err := router.Run(":" + config.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
