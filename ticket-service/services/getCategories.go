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

type CategoryService struct {
	db      *mongo.Collection
	sfGroup singleflight.Group
}
type ICategoryService interface {
	GetCategoryByEventID(context.Context, string) ([]models.CategoryResponse, error)
}

func NewCategoryService(db *mongo.Collection) ICategoryService {
	return &CategoryService{db: db}
}
func (s *CategoryService) GetCategoryByEventID(ctx context.Context, id string) ([]models.CategoryResponse, error) {
	res, err, _ := s.sfGroup.Do(id, func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		db := s.db
		pipeline := mongo.Pipeline{
			{{Key: "$match", Value: bson.D{{Key: "eventId", Value: id}}}},

			{{Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$category"},
				{Key: "price", Value: bson.D{{Key: "$first", Value: "$price"}}},
				{Key: "total_tickets", Value: bson.D{{Key: "$sum", Value: 1}}},
				{Key: "available_tickets", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{{Key: "$eq", Value: bson.A{"$status", "available"}}},
							1,
							0,
						}},
					}},
				}},
				{Key: "booked_tickets", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{{Key: "$ne", Value: bson.A{"$status", "available"}}},
							1,
							0,
						}},
					}},
				}},
			}}},

			{{Key: "$project", Value: bson.D{
				{Key: "name", Value: "$_id"},
				{Key: "price", Value: 1},
				{Key: "capacity", Value: "$total_tickets"},
				{Key: "available_tickets", Value: 1},
				{Key: "booked_tickets", Value: 1},
			}}},

			{{Key: "$sort", Value: bson.D{{Key: "price", Value: -1}}}},
		}

		cursor, err := db.Aggregate(ctx, pipeline)
		if err != nil {
			return nil, fmt.Errorf("aggregation failed: %w", err)
		}
		defer cursor.Close(ctx)

		var results []models.CategoryResponse
		if err := cursor.All(ctx, &results); err != nil {
			return nil, fmt.Errorf("failed to decode results: %w", err)
		}

		return results, nil
	})

	if err != nil {
		return nil, err
	}

	return res.([]models.CategoryResponse), nil
}
