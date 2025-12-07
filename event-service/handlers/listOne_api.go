package handlers

import (
	"github.com/devashish0812/event-service/services"

	// "net/http"
	// "time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type ListOneEventHandler struct {
	service services.GetOneEventService
}

func NewListOneEventHandler(service services.GetOneEventService) *ListOneEventHandler {
	return &ListOneEventHandler{service: service}
}
func (h *ListOneEventHandler) ListOneEvent(c *gin.Context) {

	// id := c.Query("id")
	id := c.Param("id")
	event, err := h.service.GetOneEvent(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, event)
}
