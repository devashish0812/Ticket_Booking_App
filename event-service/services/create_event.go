package services

import (
	"context"
	"time"

	"event-service/config"
	"event-service/models"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type EventService interface {
	CreateEvent(ctx context.Context, user models.Event) error
}

type eventService struct {
	con *config.MongoConfig
}

func NewEventService(con *config.MongoConfig) EventService {
	return &eventService{con: con}
}

func (s *eventService) CreateEvent(ctx context.Context, event models.Event) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := s.con.EventCol.InsertOne(ctx, event)
	return err
}
