package handlers

import (
	"github.com/devashish0812/event-service/models"
	"github.com/devashish0812/event-service/services"
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
	var event models.Master

	role, exists := c.Get("role")
	if !exists || role != "organizer" {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateEvent(c.Request.Context(), event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "event created successfully"})

}
