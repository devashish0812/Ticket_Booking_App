package handlers

import (
	"event-service/models"
	"event-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service services.EventService
}

func NewEventHandler(service services.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var user models.Event

	role, exists := c.Get("role")
	if !exists || role != "organizer" {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateEvent(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "event created successfully"})

}
