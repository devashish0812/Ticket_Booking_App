package routes

import (
	"api-gateway/config"
	"api-gateway/proxy"

	"github.com/gin-gonic/gin"
)

func RegisterTicketsRoutes(r *gin.Engine, cfg config.ServiceConfig) {
	auth := r.Group("/tickets")
	{
		auth.GET("/categories/:id", func(c *gin.Context) {
			id := c.Param("id")
			targetPath := "/tickets/categories/" + id

			proxy.Forward(cfg.TicketService, targetPath)(c)
		})
		auth.GET("/events/:eventId/categories/:categoryName", func(c *gin.Context) {
			eventId := c.Param("eventId")
			categoryName := c.Param("categoryName")
			targetPath := "/tickets/events/" + eventId + "/categories/" + categoryName

			proxy.Forward(cfg.TicketService, targetPath)(c)
		})
	}
}
