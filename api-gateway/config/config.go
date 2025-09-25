package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServiceConfig struct {
	AuthService  string
	EventService string
}

func LoadConfig() ServiceConfig {
	_ = godotenv.Load() // load from .env (optional in prod)

	authURL := os.Getenv("AUTH_SERVICE_URL")
	if authURL == "" {
		log.Fatal("AUTH_SERVICE_URL not set")
	}

	eventURL := os.Getenv("EVENT_SERVICE_URL")
	if eventURL == "" {
		log.Fatal("EVENT_SERVICE_URL not set")
	}

	return ServiceConfig{
		AuthService:  authURL,
		EventService: eventURL,
	}
}
