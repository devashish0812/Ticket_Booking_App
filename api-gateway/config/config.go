package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServiceConfig struct {
	AuthService   string
	EventService  string
	TicketService string
}

func LoadConfig() ServiceConfig {
	_ = godotenv.Load()
	authURL := os.Getenv("AUTH_SERVICE_URL")
	if authURL == "" {
		log.Fatal("AUTH_SERVICE_URL not set")
	}

	eventURL := os.Getenv("EVENT_SERVICE_URL")
	if eventURL == "" {
		log.Fatal("EVENT_SERVICE_URL not set")
	}

	ticketURL := os.Getenv("TICKET_SERVICE_URL")
	if eventURL == "" {
		log.Fatal("TICKET_SERVICE_URL not set")
	}
	return ServiceConfig{
		AuthService:   authURL,
		EventService:  eventURL,
		TicketService: ticketURL,
	}
}
