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
	mongoCfg := config.InitMongo()

	createTicketService := services.NewTicketService(mongoCfg.TicketCol)

	outboxService := services.NewWorker("", , createTicketService.(*services.TicketService))
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		log.Println("Outbox worker started...")
		outboxService.Run(ctx)
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

}
