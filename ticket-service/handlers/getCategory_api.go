package handlers

import (
	"log"
	"ticket-service/services"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}
func (h *CategoryHandler) ListAllCategories(c *gin.Context) {
	id := c.Param("id")
	categories, err := h.service.GetCategoryByEventID(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.Print(categories)
	c.JSON(200, categories)
}
