package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/devashish0812/event-service/config"
	"github.com/devashish0812/event-service/handlers"
	"github.com/devashish0812/event-service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	mongoCfg := config.InitMongo()

	eventService := services.NewEventService(mongoCfg)
	eventHandler := handlers.NewEventHandler(eventService)

	listAllEventService := services.NewGetAllEventService(mongoCfg)
	listAllEventHandler := handlers.NewListAllEventHandler(listAllEventService)

	listOneEventService := services.NewGetOneEventService(mongoCfg)
	listOneEventHandler := handlers.NewListOneEventHandler(listOneEventService)

	listAllEventForOrgService := services.NewGetAllEventForOrgService(mongoCfg)
	listAllEventForOrgHandler := handlers.NewListAllEventForOrgHandler(listAllEventForOrgService)

	authMiddleware := handlers.NewAuthMiddleware(mongoCfg.JWTSecret)

	r := gin.Default()
	events := r.Group("/events", authMiddleware.RequireAuth())
	{
		events.POST("/create", eventHandler.CreateEvent)
		events.GET("/getall", listAllEventHandler.ListEvents)
		events.GET("/get/:id", listOneEventHandler.ListOneEvent)
		events.GET("/getallForOrg", listAllEventForOrgHandler.ListEventsForOrg)
	}

	outboxService := services.NewOutboxService(mongoCfg)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		log.Println("Outbox worker started...")
		outboxService.StartWorker(ctx, "worker-1")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutdown signal received, stopping outbox worker...")
		cancel()
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Event Service running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
