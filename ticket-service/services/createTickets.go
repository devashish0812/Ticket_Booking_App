package services

import (
	"context"
	"fmt"
	"time"

	"ticket-service/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketService struct {
	db *mongo.Collection
}

type ITicketService interface {
	CreateTicket([]models.Ticket) error
}

func NewTicketService(db *mongo.Collection) ITicketService {
	return &TicketService{db: db}
}
func (s *TicketService) CreateTicket(Tickets []models.Ticket) error {
	now := time.Now()

	for _, ticket := range Tickets {
		for i := 0; i < ticket.Quantity; i++ {
			seat := models.Seat{
				ID:         primitive.NewObjectID(),
				EventID:    ticket.EventId,
				Category:   ticket.Type,
				SeatNumber: string(ticket.Type[0]) + fmt.Sprintf("-%d", i+1),
				Price:      ticket.Price,
				Status:     "available",
				CreatedAt:  now,
				UpdatedAt:  now,
			}

			_, err := s.db.InsertOne(context.Background(), seat)

			if err != nil {
				return fmt.Errorf("failed to create ticket: %w", err)
			}
		}

	}

	return nil
}

// here ticket means the general ticket info like price,type,quantity etc
// seats will be created in seat collection
