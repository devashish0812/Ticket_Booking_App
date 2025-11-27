package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"ticket-service/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TicketService struct {
	db *mongo.Collection
}

type ITicketService interface {
	CreateTicket(context.Context, []models.Ticket) error
}

func NewTicketService(db *mongo.Collection) ITicketService {
	return &TicketService{db: db}
}

func (s *TicketService) CreateTicket(ctx context.Context, Tickets []models.Ticket) error {
	now := time.Now()
	var allSeats []interface{}

	for _, ticket := range Tickets {
		for i := 0; i < ticket.Quantity; i++ {
			seat := models.Seat{
				ID:         primitive.NewObjectID(),
				EventID:    ticket.EventId,
				Category:   ticket.Type,
				SeatNumber: ticket.Type + "-" + fmt.Sprintf("%04d", i+1),
				Price:      ticket.Price,
				Status:     "available",
				CreatedAt:  now,
				UpdatedAt:  now,
			}
			allSeats = append(allSeats, seat)
		}
	}

	if len(allSeats) == 0 {
		return nil
	}

	batchSize := 1000
	for i := 0; i < len(allSeats); i += batchSize {
		end := i + batchSize
		if end > len(allSeats) {
			end = len(allSeats)
		}

		batch := allSeats[i:end]

		insertCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		opts := options.InsertMany().SetOrdered(false)

		result, err := s.db.InsertMany(insertCtx, batch, opts)
		cancel()

		if err != nil {
			if result != nil {
				log.Printf("Partial insert: %d/%d seats created",
					len(result.InsertedIDs), len(batch))
			}
			return fmt.Errorf("failed to insert batch %d-%d: %w", i, end, err)
		}

		log.Printf("Successfully inserted batch %d-%d (%d seats)", i, end, len(batch))
	}

	return nil
}

// here ticket means the general ticket info like price,type,quantity etc
// seats will be created in seat collection
