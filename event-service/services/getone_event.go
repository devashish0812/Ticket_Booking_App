package services

import (
	"context"
	"time"

	"github.com/devashish0812/event-service/config"
	"github.com/devashish0812/event-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type GetOneEventService interface {
	GetOneEvent(ctx context.Context, id string) (models.Event, error)
}

type getOneeventService struct {
	con *config.MongoConfig
}

func NewGetOneEventService(con *config.MongoConfig) GetOneEventService {
	return &getOneeventService{con: con}
}

func (s *getOneeventService) GetOneEvent(ctx context.Context, id string) (models.Event, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Event{}, err
	}

	collection := s.con.EventCol
	var event models.Event

	err = collection.FindOne(timeoutCtx, bson.M{"_id": objID}).Decode(&event)
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}
