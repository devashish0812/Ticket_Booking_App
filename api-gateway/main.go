package main

import (
	"api-gateway/config"
	"api-gateway/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()
	allowedOrigins := map[string]bool{
		"https://ticket-booking-app-xi.vercel.app": true,
		"http://localhost:5173":                    true,
		"http://localhost:3000":                    true,
		"https://user-frontend-kappa.vercel.app":   true,
	}
	r.GET("/health", func(c *gin.Context) { // for pinging the service health
		c.JSON(200, gin.H{
			"status":  "alive",
			"service": "api-gateway",
		})
	})
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
	routes.RegisterTicketsRoutes(r, cfg)

	// Run Gateway
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default for local development
	}
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal(err)
	}
}
