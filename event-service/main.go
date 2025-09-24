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
	eventService := services.NewEventService(mongoCfg)
	eventHandler := handlers.NewEventHandler(eventService)

	listAllEventService := services.NewGetAllEventService(mongoCfg)
	listAllEventHandler := handlers.NewListAllEventHandler(listAllEventService)

	listOneEventService := services.NewGetOneEventService(mongoCfg)
	listOneEventHandler := handlers.NewListOneEventHandler(listOneEventService)

	listAllEventForOrgService := services.NewGetAllEventForOrgService(mongoCfg)
	listAllEventForOrgHandler := handlers.NewListAllEventForOrgHandler(listAllEventForOrgService)

	authMiddleware := handlers.NewAuthMiddleware(mongoCfg.JWTSecret)
	// 3) Routes
	r := gin.Default()

	events := r.Group("/events", authMiddleware.RequireAuth())
	{
		events.POST("/create", eventHandler.CreateEvent)
		events.GET("/getall", listAllEventHandler.ListEvents)
		events.GET("/get", listOneEventHandler.ListOneEvent)
		events.GET("/getallForOrg", listAllEventForOrgHandler.ListEventsForOrg)
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
