package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("  No .env file found, relying on system environment variables")
	}
}

func GetMongoURI() string {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal(" MONGO_URI is not set in environment variables")
	}
	return uri
}
