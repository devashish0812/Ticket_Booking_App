package handlers

import (
	"github.com/devashish0812/event-service/models"
	"github.com/devashish0812/event-service/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListAllEventHandler struct {
	service services.GetAllEventService
}

func NewListAllEventHandler(service services.GetAllEventService) *ListAllEventHandler {
	return &ListAllEventHandler{service: service}
}
func (h *ListAllEventHandler) ListEvents(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	// Wrap into a struct
	filterReq := models.EventFilterRequest{
		Category: c.Query("category"),
		Date:     c.Query("date"),
		SortBy:   c.DefaultQuery("sortBy", "startDateTime"),
		Order:    c.DefaultQuery("order", "asc"),
		Page:     page,
		Limit:    limit,
	}

	events, totalCount, err := h.service.GetAllEvent(c.Request.Context(), filterReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"events":     events,
		"totalCount": totalCount,
	})
}
