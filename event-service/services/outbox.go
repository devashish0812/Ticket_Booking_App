package services

import (
	"context"
	"log"
	"time"

	"github.com/devashish0812/ticketingtool/common/kafka"

	"github.com/devashish0812/event-service/config"
	"github.com/devashish0812/event-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

type OutboxService interface {
	StartWorker(ctx context.Context, id string)
}

type outboxService struct {
	con *config.MongoConfig
}

func NewOutboxService(con *config.MongoConfig) OutboxService {
	return &outboxService{con: con}
}

func (s *outboxService) StartWorker(ctx context.Context, id string) {
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Printf("Outbox worker %s stopped: %v", id, ctx.Err())
				return
			case <-ticker.C:
				events, err := s.findPendingEvents(ctx)
				// event here means outbox events
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
	}()
}

func (s *outboxService) findPendingEvents(ctx context.Context) ([]models.OutboxEvent, error) {
	// Mongo query: find documents where processed = false
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
	// Mongo update: set status = "published"
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := s.con.OutboxCol

	_, err := collection.UpdateOne(timeoutCtx, bson.M{"_id": id}, bson.M{"$set": bson.M{"processed": true}})
	if err != nil {
		return err
	}
	return nil
}
func (s *outboxService) publishWithRetry(event models.OutboxEvent) error {
	producer := kafka.NewProducer(&kafka.Config{
		Brokers: []string{"localhost:9092"},
	})
	defer producer.Close()
	message := struct {
		EventID string      `json:"eventId"`
		Payload interface{} `json:"payload"`
	}{
		EventID: event.EventID,
		Payload: event.Payload,
	}

	err := producer.Publish("ticketDetails.created", message)
	if err != nil {
		log.Println("Outbox publish failed, will retry later:", err)
		return err
	}
	return nil
}
