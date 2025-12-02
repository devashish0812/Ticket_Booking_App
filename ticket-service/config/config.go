package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceDependencies struct {
	Client    *mongo.Client
	TicketCol *mongo.Collection
	JWTSecret string
	Topic     string
	GroupID   string
}

func LoadDependencies() *ServiceDependencies {
	uri := os.Getenv("MONGO_URI") // from .env
	if uri == "" {
		log.Fatal("MONGO_URI not set in environment")
	}

	// connect
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Pick DB and Collection
	db := client.Database("ticketingtool")
	col := db.Collection("Tickets")
	JWTSecret := os.Getenv("JWT_SECRET")
	Topic := os.Getenv("KAFKA_TOPIC")
	GroupID := os.Getenv("KAFKA_GROUP_ID")

	return &ServiceDependencies{
		Client:    client,
		TicketCol: col,
		JWTSecret: JWTSecret,
		Topic:     Topic,
		GroupID:   GroupID,
	}
}
