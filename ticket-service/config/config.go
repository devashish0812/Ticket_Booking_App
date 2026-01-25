package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceDependencies struct {
	MongoClient *mongo.Client
	TicketCol   *mongo.Collection
	RedisClient *redis.Client
	JWTSecret   string
	Topic       string
	GroupID     string
}

func LoadDependencies() *ServiceDependencies {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI not set in environment")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Mongo Connection Error:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo Ping Failed:", err)
	}

	db := client.Database("ticketingtool")
	col := db.Collection("Tickets")

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("REDIS_URL not set in environment")
	}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatal("Invalid Redis URL:", err)
	}

	rdb := redis.NewClient(opt)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Redis Connection Failed:", err)
	}

	return &ServiceDependencies{
		MongoClient: client,
		TicketCol:   col,
		RedisClient: rdb,
		JWTSecret:   os.Getenv("JWT_SECRET"),
		Topic:       os.Getenv("KAFKA_TOPIC"),
		GroupID:     os.Getenv("KAFKA_GROUP_ID"),
	}
}
