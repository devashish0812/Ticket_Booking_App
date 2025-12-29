package handlers

import (
	"log"
	"ticket-service/services"

	"github.com/gin-gonic/gin"
)

type SectionHandler struct {
	service *services.SectionService
}

func NewSectionHandler(service *services.SectionService) *SectionHandler {
	return &SectionHandler{service: service}
}
func (h *SectionHandler) ListAllSection(c *gin.Context) {
	eventID := c.Param("eventId")
	categoryName := c.Param("category")
	sections, err := h.service.GetSectionByEventID(c.Request.Context(), eventID, categoryName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.Print(sections)
	c.JSON(200, sections)
}
