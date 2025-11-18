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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://ticket-booking-app-xi.vercel.app", "http://localhost:5173", "http://localhost:3000", "https://user-frontend-kappa.vercel.app"}, // React dev server
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Register routes
	routes.RegisterAuthRoutes(r, cfg)
	routes.RegisterDashboardRoutes(r, cfg)

	// Run Gateway
	r.Run(":8081") // Gateway will run on port 8081
}
