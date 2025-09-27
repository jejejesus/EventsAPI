package handlers

import (
	"EventsAPI/internal/domain/entities"
	"EventsAPI/internal/usecases"
	"fmt"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventUseCase *usecases.EventUseCase
}

func NewEventHandler(eventUseCase *usecases.EventUseCase) *EventHandler {
	return &EventHandler{eventUseCase: eventUseCase}
}

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event with title, description, date, and location
// @Tags events
// @Accept json
// @Produce json
// @Param event body entities.EventRequest true "Event creation data"
// @Success 201 {object} entities.EventResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /events [post]
// @Security Bearer
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var req entities.EventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	newEvent := &entities.Event{
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		DateTime:    req.DateTime,
		MaxCapacity: req.MaxCapacity,
		UserID:      userID.(uint),
	}

	err := h.eventUseCase.CreateEvent(c.Request.Context(), newEvent)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "Event created successfully",
	})
}

// ListEvents godoc
// @Summary List all events
// @Description Retrieve a list of all events
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {array} entities.EventResponse
// @Failure 500 {object} map[string]string
// @Router /events [get]
func (h *EventHandler) ListEvents(c *gin.Context) {
	events, err := h.eventUseCase.ListEvents(c.Request.Context(), 100, 0)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var response []entities.EventResponse
	for _, event := range events {
		response = append(response, entities.EventResponse{
			ID:             event.ID,
			Title:          event.Title,
			Description:    event.Description,
			Location:       event.Location,
			DateTime:       event.DateTime,
			MaxCapacity:    event.MaxCapacity,
			UserID:         event.UserID,
			AttendeesCount: len(event.Attendees),
			CreatedAt:      event.CreatedAt,
		})
	}

	c.JSON(200, response)
}

// GetEvent godoc
// @Summary Get event by ID
// @Description Retrieve a specific event by its ID
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} entities.EventResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /events/{id} [get]
func (h *EventHandler) GetEvent(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := h.eventUseCase.GetEventByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	response := entities.EventResponse{
		ID:             event.ID,
		Title:          event.Title,
		Description:    event.Description,
		Location:       event.Location,
		DateTime:       event.DateTime,
		MaxCapacity:    event.MaxCapacity,
		UserID:         event.UserID,
		AttendeesCount: len(event.Attendees),
		CreatedAt:      event.CreatedAt,
	}

	c.JSON(200, response)
}

// UpdateEvent godoc
// @Summary Update an event
// @Description Update an existing event by its ID
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Param event body entities.EventRequest true "Event update data"
// @Success 200 {object} entities.EventResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /events/{id} [put]
// @Security Bearer
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	var req entities.EventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	event, err := h.eventUseCase.GetEventByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	if event.UserID != userID.(uint) {
		c.JSON(401, gin.H{"error": "You are not the owner of this event"})
		return
	}

	event.Title = req.Title
	event.Description = req.Description
	event.Location = req.Location
	event.DateTime = req.DateTime
	event.MaxCapacity = req.MaxCapacity

	err = h.eventUseCase.UpdateEvent(c.Request.Context(), event)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := entities.EventResponse{
		ID:             event.ID,
		Title:          event.Title,
		Description:    event.Description,
		Location:       event.Location,
		DateTime:       event.DateTime,
		MaxCapacity:    event.MaxCapacity,
		UserID:         event.UserID,
		AttendeesCount: len(event.Attendees),
		CreatedAt:      event.CreatedAt,
	}

	c.JSON(200, response)
}

// DeleteEvent godoc
// @Summary Delete an event
// @Description Delete an existing event by its ID
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /events/{id} [delete]
// @Security Bearer
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	event, err := h.eventUseCase.GetEventByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	if event.UserID != userID.(uint) {
		c.JSON(401, gin.H{"error": "You are not the owner of this event"})
		return
	}

	err = h.eventUseCase.DeleteEvent(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

// GetMyEvents godoc
// @Summary Get my events
// @Description Retrieve events created by the authenticated user
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {array} entities.EventResponse
// @Failure 401 {object} map[string]string
// @Router /events/my [get]
// @Security Bearer
func (h *EventHandler) GetMyEvents(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	events, err := h.eventUseCase.GetUserEvents(c.Request.Context(), userID.(uint), 100, 0)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var response []entities.EventResponse
	for _, event := range events {
		response = append(response, entities.EventResponse{
			ID:             event.ID,
			Title:          event.Title,
			Description:    event.Description,
			Location:       event.Location,
			DateTime:       event.DateTime,
			MaxCapacity:    event.MaxCapacity,
			UserID:         event.UserID,
			AttendeesCount: len(event.Attendees),
			CreatedAt:      event.CreatedAt,
		})
	}

	c.JSON(200, response)
}
