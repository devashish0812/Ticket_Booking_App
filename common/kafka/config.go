package kafka

import (
	"log"
	"os"
)

type Config struct {
	Brokers []string
}

func LoadConfig() *Config {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "localhost:9092" // default for local dev
		log.Println("⚠️  KAFKA_BROKER not set, using localhost:9092")
	}

	return &Config{
		Brokers: []string{broker},
	}
}
