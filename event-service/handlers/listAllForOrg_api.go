package handlers

import (
	"github.com/devashish0812/event-service/models"
	"github.com/devashish0812/event-service/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListAllEventForOrgHandler struct {
	service services.GetAllEventForOrgService
}

func NewListAllEventForOrgHandler(service services.GetAllEventForOrgService) *ListAllEventForOrgHandler {
	return &ListAllEventForOrgHandler{service: service}
}
func (h *ListAllEventForOrgHandler) ListEventsForOrg(c *gin.Context) {

	role, exists := c.Get("role")
	if !exists || role != "organizer" {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	// Wrap into a struct
	filterReq := models.EventFilterRequest{
		Type:  c.Query("type"),
		Page:  page,
		Limit: limit,
	}

	events, err := h.service.GetAllEventForOrg(c.Request.Context(), filterReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, events)
}
