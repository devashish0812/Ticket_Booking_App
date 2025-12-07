package routes

import (
	"api-gateway/config"
	"api-gateway/proxy"

	"github.com/gin-gonic/gin"
)

func RegisterEventsRoutes(r *gin.Engine, cfg config.ServiceConfig) {
	auth := r.Group("/events")
	{
		auth.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			targetPath := "/events/get/" + id

			proxy.Forward(cfg.EventService, targetPath)(c)
		})
	}
}
