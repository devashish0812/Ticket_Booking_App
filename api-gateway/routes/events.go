package routes

import (
	"api-gateway/config"
	"api-gateway/proxy"

	"github.com/gin-gonic/gin"
)

func RegisterEventsRoutes(r *gin.Engine, cfg config.ServiceConfig) {
	auth := r.Group("/events")
	{
		// Gateway: GET /events/:id  →  EventService: GET /events/:id
		auth.GET("/:id", proxy.Forward(cfg.EventService, "/events/get/:id"))

	}
}
