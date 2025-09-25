package routes

import (
	"api-gateway/config"
	"api-gateway/proxy"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, cfg config.ServiceConfig) {
	auth := r.Group("/auth")
	{
		// Gateway: POST /auth/login  →  AuthService: POST /users/login
		auth.POST("/login", proxy.Forward(cfg.AuthService, "/users/login"))

		// Gateway: POST /auth/signup → AuthService: POST /users/signup
		auth.POST("/signup", proxy.Forward(cfg.AuthService, "/users/signup"))

		// Gateway: POST /auth/refresh → AuthService: POST /auth/refresh
		auth.POST("/refresh", proxy.Forward(cfg.AuthService, "/auth/refresh"))
	}
}
