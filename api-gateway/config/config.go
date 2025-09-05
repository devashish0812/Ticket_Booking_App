package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServiceConfig struct {
	AuthService string
}

func LoadConfig() ServiceConfig {
	_ = godotenv.Load() // load from .env (optional in prod)

	authURL := os.Getenv("AUTH_SERVICE_URL")
	if authURL == "" {
		log.Fatal("AUTH_SERVICE_URL not set")
	}

	return ServiceConfig{
		AuthService: authURL,
	}
}
