package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ticket-service/config"
	"ticket-service/services"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dependencies := config.LoadDependencies()

	createTicketService := services.NewTicketService(dependencies.TicketCol)

	outboxService := services.NewWorker(dependencies.Topic, dependencies.GroupID, createTicketService.(*services.TicketService))

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		log.Println("Outbox worker started...")
		outboxService.Run(ctx)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan

	log.Println("Shutdown signal received, stopping outbox worker...")
	cancel()
	time.Sleep(2 * time.Second)
	log.Println("Service stopped gracefully")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Ticket Service running on port %s\n", port)

}
