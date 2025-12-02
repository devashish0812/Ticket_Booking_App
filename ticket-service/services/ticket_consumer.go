package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ticket-service/models"
	"time"

	customkafka "github.com/devashish0812/Ticket_Booking_App/common/kafka"
	"github.com/segmentio/kafka-go"
)

type Worker struct {
	topic     string
	groupID   string
	ticketSvc *TicketService
}

func NewWorker(topic string, groupID string, ticketSvc *TicketService) *Worker {
	return &Worker{
		topic:     topic,
		groupID:   groupID,
		ticketSvc: ticketSvc,
	}
}

func (w *Worker) Run(ctx context.Context) {
	cfg := customkafka.LoadConfig()
	kConsumer := customkafka.NewConsumer(cfg, w.groupID, w.topic)

	defer func() {
		if err := kConsumer.Close(); err != nil {
			log.Printf("Error closing consumer: %v", err)
		}
	}()

	fmt.Printf("Worker started for Topic: %s | Group: %s\n", w.topic, w.groupID)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		msg, err := kConsumer.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Printf("Error reading message: %v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}

		if err := w.processMessage(ctx, msg); err != nil {
			log.Printf("Error processing message: %v\n", err)
			// Handle failure (retry, DLQ, etc.)
		}
	}
}

func (w *Worker) processMessage(ctx context.Context, msg kafka.Message) error {

	type EventWrapper struct {
		EventID string          `json:"eventId"`
		Payload []models.Ticket `json:"payload"`
	}

	var wrapper EventWrapper

	if err := json.Unmarshal(msg.Value, &wrapper); err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}
	if err := w.ticketSvc.CreateTicket(ctx, wrapper.Payload); err != nil {
		return fmt.Errorf("create tickets error: %w", err)
	}

	return nil
}
