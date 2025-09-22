package main

import (
	"log"
	"os"

	"event-service/config"
	"event-service/handlers"
	"event-service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1) Initialize Mongo via config (reads MONGO_URI from env)
	_ = godotenv.Load()
	mongoCfg := config.InitMongo()
	// 2) Wire layers
	eventService := services.NewEventService(mongoCfg)     // signup service
	eventHandler := handlers.NewEventHandler(eventService) // signup handler

	// 3) Routes
	r := gin.Default()

	events := r.Group("/events")
	{
		events.POST("/create", eventHandler.CreateEvent)
	}

	//r.POST("/auth/refresh", authMiddleware.RequireAuth(), authHandler.GetRefreshToken)

	// 4) Start
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
