package handlers

import (
	"EventsAPI/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AttendeeHandler struct {
	attendeeUseCase *usecases.AttendeeUseCase
}

func NewAttendeeHandler(attendeeUseCase *usecases.AttendeeUseCase) *AttendeeHandler {
	return &AttendeeHandler{attendeeUseCase: attendeeUseCase}
}

// RegisterForEvent godoc
// @Summary Register for an event
// @Description Register the authenticated user for a specific event
// @Tags attendees
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /attendees/event/{event_id}/register [post]

func (h *AttendeeHandler) RegisterForEvent(c *gin.Context) {
	eventIDStr := c.Param("event_id")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Convert eventID from string to uint
	eventIDUint, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	err = h.attendeeUseCase.RegisterForEvent(c.Request.Context(), userID.(uint), uint(eventIDUint))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Registered for event successfully"})
	c.JSON(200, gin.H{"message": "Registered for event successfully"})
}

// UnregisterFromEvent godoc
// @Summary Unregister from an event
// @Description Unregister the authenticated user from a specific event
// @Tags attendees
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /attendees/event/{event_id}/unregister [post]
func (h *AttendeeHandler) UnregisterFromEvent(c *gin.Context) {
	eventIDStr := c.Param("event_id")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Convert eventID from string to uint
	eventIDUint, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	err = h.attendeeUseCase.UnregisterFromEvent(c.Request.Context(), userID.(uint), uint(eventIDUint))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Unregistered from event successfully"})
}

// GetMyRegistrations godoc
// @Summary Get my event registrations
// @Description Retrieve a list of events the authenticated user is registered for
// @Tags attendees
// @Accept json
// @Produce json
// @Success 200 {array} entities.EventResponse
// @Failure 401 {object} map[string]string
// @Router /attendees/my [get]
func (h *AttendeeHandler) GetMyRegistrations(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	registrations, err := h.attendeeUseCase.GetMyRegistrations(c.Request.Context(), userID.(uint), 10, 0)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, registrations)
}

// GetEventAttendees godoc
// @Summary Get event attendees
// @Description Retrieve a list of users registered for a specific event
// @Tags attendees
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Success 200 {array} entities.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /attendees/event/{event_id}/attendees [get]
func (h *AttendeeHandler) GetEventAttendees(c *gin.Context) {
	eventIDStr := c.Param("event_id")
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Convert eventID from string to uint
	eventIDUint, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	attendees, err := h.attendeeUseCase.GetEventAttendees(c.Request.Context(), uint(eventIDUint), 10, 0)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, attendees)
}
