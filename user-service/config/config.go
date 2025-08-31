package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Client    *mongo.Client
	UserCol   *mongo.Collection
	TokenCol  *mongo.Collection
	JWTSecret string
}

func InitMongo() *MongoConfig {
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
	col := db.Collection("users")
	token := db.Collection("refresh_tokens")
	JWTSecret := os.Getenv("JWT_SECRET")
	return &MongoConfig{
		Client:    client,
		UserCol:   col,
		JWTSecret: JWTSecret,
		TokenCol:  token,
	}
}
