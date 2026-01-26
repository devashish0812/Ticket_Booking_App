package handlers

import (
	"ticket-service/services"

	"github.com/gin-gonic/gin"
)

type SeatsHandler struct {
	service *services.SeatsService
}

func NewSeatsHandler(service *services.SeatsService) *SeatsHandler {
	return &SeatsHandler{service: service}
}
func (h *SeatsHandler) ListAllSeats(c *gin.Context) {
	eventID := c.Param("eventId")
	sectionName := c.Param("section")
	seats, err := h.service.GetSeats(c.Request.Context(), eventID, sectionName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, seats)
}
