package services

import (
	"context"
	"log"
	"time"

	"github.com/devashish0812/Ticket_Booking_App/common/kafka"

	"github.com/devashish0812/event-service/config"
	"github.com/devashish0812/event-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

type OutboxService interface {
	StartWorker(ctx context.Context, id string)
	Close() error
}

type outboxService struct {
	con      *config.MongoConfig
	producer *kafka.Producer
	topic    string
}

func NewOutboxService(con *config.MongoConfig, topic string) OutboxService {
	cfg := kafka.LoadConfig()
	producer := kafka.NewProducer(cfg)

	return &outboxService{
		con:      con,
		producer: producer,
		topic:    topic,
	}
}

func (s *outboxService) Close() error {
	return s.producer.Close()
}

func (s *outboxService) StartWorker(ctx context.Context, id string) {

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	log.Printf("Outbox worker %s started", id)
	for {
		select {
		case <-ctx.Done():
			log.Printf("Outbox worker %s stopped: %v", id, ctx.Err())
			return
		case <-ticker.C:
			events, err := s.findPendingEvents(ctx)
			if err != nil {
				log.Println("Error reading outbox:", err)
				continue
			}

			for _, evt := range events {
				select {
				case <-ctx.Done():
					log.Println("Context cancelled while processing event batch, stopping.")
					return
				default:
				}

				if err := s.publishWithRetry(evt); err != nil {
					log.Println("Publish failed:", err)
					continue
				}
				if err := s.markAsPublished(ctx, evt.EventID); err != nil {
					log.Printf("Failed to mark event %s as published: %v", evt.ID, err)
				}

			}
		}
	}

}

func (s *outboxService) findPendingEvents(ctx context.Context) ([]models.OutboxEvent, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := s.con.OutboxCol
	var events []models.OutboxEvent

	cursor, err := collection.Find(timeoutCtx, bson.M{"processed": false})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}

func (s *outboxService) markAsPublished(ctx context.Context, id string) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := s.con.OutboxCol

	_, err := collection.UpdateOne(timeoutCtx, bson.M{"eventId": id}, bson.M{"$set": bson.M{"processed": true}})
	if err != nil {
		log.Println("Error while Updating Status in the DB", err)
		return err
	}
	log.Println("Event Status Updated in the DB", id)
	return nil
}

func (s *outboxService) publishWithRetry(event models.OutboxEvent) error {
	message := struct {
		EventID string          `json:"eventId"`
		Payload []models.Ticket `json:"payload"`
	}{
		EventID: event.EventID,
		Payload: event.Payload,
	}

	err := s.producer.Publish(s.topic, message)
	if err != nil {
		log.Println("Outbox publish failed, will retry later:", err)
		return err
	}

	return nil
}
