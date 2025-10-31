package services

import (
	"context"
	"errors"
	"time"

	"github.com/devashish0812/event-service/config"
	"github.com/devashish0812/event-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type EventService interface {
	CreateEvent(ctx context.Context, allDetails models.Master) error
}

type eventService struct {
	con *config.MongoConfig
}

func NewEventService(con *config.MongoConfig) EventService {
	return &eventService{con: con}
}

func (s *eventService) CreateEvent(ctx context.Context, allDetails models.Master) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	event := allDetails.Event
	ticket := allDetails.Tickets
	res, err := s.con.EventCol.InsertOne(ctx, event)
	if err != nil {
		return err
	}
	eventID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return errors.New("failed to retrieve Event ID")
	}

	// Set EventId for all tickets
	for i := range ticket {
		ticket[i].EventId = eventID.Hex()
	}

	outboxRecord := bson.M{
		"eventId":   eventID,
		"eventType": "TicketsCreated",
		"payload":   ticket,
		"createdAt": time.Now(),
		"processed": false,
	}

	_, err = s.con.OutboxCol.InsertOne(ctx, outboxRecord)

	if err != nil {
		return err
	}

	return err
}
