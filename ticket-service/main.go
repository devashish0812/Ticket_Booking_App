package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ticket-service/config"
	"ticket-service/handlers"
	"ticket-service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dependencies := config.LoadDependencies()

	createTicketService := services.NewTicketService(dependencies.TicketCol)
	outboxService := services.NewWorker(dependencies.Topic, dependencies.GroupID, createTicketService.(*services.TicketService))

	categoryService := services.NewCategoryService(dependencies.TicketCol)
	categoryHandler := handlers.NewCategoryHandler(categoryService.(*services.CategoryService))

	sectionService := services.NewSectionService(dependencies.TicketCol)
	sectionHandler := handlers.NewSectionHandler(sectionService.(*services.SectionService))

	seatService := services.NewSeatLockService(dependencies.RedisClient)
	seatHandler := handlers.NewSeatLockHandler(seatService)

	authMiddleware := handlers.NewAuthMiddleware(dependencies.JWTSecret)

	r := gin.Default()
	tickets := r.Group("/tickets", authMiddleware.RequireAuth())
	{
		tickets.GET("/categories/:id", categoryHandler.ListAllCategories)
		tickets.GET("/events/:eventId/categories/:category", sectionHandler.ListAllSection)
		tickets.POST("seats/lock", seatHandler.HandleLockSeat)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		log.Println("Outbox worker started...")
		outboxService.Run(ctx)
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	go func() {
		log.Printf("Ticket Service running on port %s\n", port)
		if err := r.Run(":" + port); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutdown signal received, stopping outbox worker...")
	cancel()
	time.Sleep(2 * time.Second)
	log.Println("Service stopped gracefully")
}
