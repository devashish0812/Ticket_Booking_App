package services

import (
	"context"
	"ticket-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SeatsService struct {
	db *mongo.Collection
}
type ISeatsService interface {
	GetSeats(ctx context.Context, eventID string, categoryName string, sectionName string) ([]models.Seat, error)
}

func NewSeatsService(db *mongo.Collection) ISeatsService {
	return &SeatsService{db: db}
}
func (s *SeatsService) GetSeats(ctx context.Context, eventID string, categoryName string, sectionName string) ([]models.Seat, error) {
	filter := bson.M{
		"eventId":  eventID,
		"category": categoryName,
		"section":  sectionName,
	}
	cursor, err := s.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var seats []models.Seat
	if err := cursor.All(ctx, &seats); err != nil {
		return nil, err
	}
	return seats, nil
}
