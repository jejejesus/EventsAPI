package routes

import (
	"EventsAPI/internal/config"
	"EventsAPI/internal/delivery/http/handlers"
	"EventsAPI/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(
	config *config.Config,
	authHandler *handlers.AuthHandler,
	eventHandler *handlers.EventHandler,
	attendeeHandler *handlers.AttendeeHandler,
	healthHandler *handlers.HealthHandler,
) *gin.Engine {

	if config.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// CORS middleware (simple version)
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := router.Group("/api/v1")

	// Health check
	api.GET("/health", healthHandler.HealthCheck)

	// Auth routes (public)
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(config))
	{
		// Events routes
		events := protected.Group("/events")
		{
			events.POST("", eventHandler.CreateEvent)
			events.GET("", eventHandler.ListEvents)
			events.GET("/my", eventHandler.GetMyEvents)
			events.GET("/:id", eventHandler.GetEvent)
			events.PUT("/:id", eventHandler.UpdateEvent)
			events.DELETE("/:id", eventHandler.DeleteEvent)
		}

		// Attendees routes
		attendees := protected.Group("/attendees")
		{
			attendees.POST("/register/:eventId", attendeeHandler.RegisterForEvent)
			attendees.POST("/unregister/:eventId", attendeeHandler.UnregisterFromEvent)
			attendees.GET("/my", attendeeHandler.GetMyRegistrations)
			attendees.GET("/event/:eventId", attendeeHandler.GetEventAttendees)
		}
	}

	return router
}
