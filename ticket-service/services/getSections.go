package services

import (
	"context"
	"fmt"
	"ticket-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/singleflight"
)

type SectionService struct {
	db      *mongo.Collection
	sfGroup singleflight.Group
}
type ISectionService interface {
	GetSectionByEventID(context.Context, string, string) ([]models.SectionResponse, error)
}

func NewSectionService(db *mongo.Collection) ISectionService {
	return &SectionService{db: db}
}
func (s *SectionService) GetSectionByEventID(ctx context.Context, eventID string, categoryName string) ([]models.SectionResponse, error) {
	sfKey := fmt.Sprintf("%s_%s", eventID, categoryName)

	res, err, _ := s.sfGroup.Do(sfKey, func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		db := s.db
		pipeline := mongo.Pipeline{
			{{Key: "$match", Value: bson.D{
				{Key: "eventId", Value: eventID},
				{Key: "category", Value: categoryName},
			}}},

			{{Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$section"},
				{Key: "total_seats", Value: bson.D{{Key: "$sum", Value: 1}}},
				{Key: "available_seats", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{{Key: "$eq", Value: bson.A{"$status", "available"}}},
							1,
							0,
						}},
					}},
				}},
			}}},

			{{Key: "$project", Value: bson.D{
				{Key: "section_name", Value: "$_id"},
				{Key: "capacity", Value: "$total_seats"},
				{Key: "available_tickets", Value: "$available_seats"},
			}}},

			{{Key: "$sort", Value: bson.D{{Key: "section_name", Value: 1}}}},
		}

		cursor, err := db.Aggregate(ctx, pipeline)
		if err != nil {
			return nil, fmt.Errorf("aggregation failed: %w", err)
		}
		defer cursor.Close(ctx)

		var results []models.SectionResponse
		if err := cursor.All(ctx, &results); err != nil {
			return nil, fmt.Errorf("failed to decode results: %w", err)
		}

		return results, nil
	})

	if err != nil {
		return nil, err
	}

	return res.([]models.SectionResponse), nil
}
