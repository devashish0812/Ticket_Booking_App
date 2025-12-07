package main

import (
	"api-gateway/config"
	"api-gateway/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load service config (.env)
	cfg := config.LoadConfig()

	// Init Gin
	r := gin.Default()
	// Define your list of allowed origins once
	allowedOrigins := map[string]bool{
		"https://ticket-booking-app-xi.vercel.app": true,
		"http://localhost:5173":                    true,
		"http://localhost:3000":                    true,
		"https://user-frontend-kappa.vercel.app":   true,
	}

	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			_, ok := allowedOrigins[origin]
			return ok
		},

		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Register routes
	routes.RegisterAuthRoutes(r, cfg)
	routes.RegisterDashboardRoutes(r, cfg)
	routes.RegisterEventsRoutes(r, cfg)

	// Run Gateway
	r.Run(":8081") // Gateway will run on port 8081
}
