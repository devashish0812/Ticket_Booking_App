package services

import (
	"context"
	"strings"
	"time"

	"event-service/config"
	"event-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type GetAllEventService interface {
	GetAllEvent(ctx context.Context, filterReq models.EventFilterRequest) ([]models.Event, error)
}

type getalleventService struct {
	con *config.MongoConfig
}

func NewGetAllEventService(con *config.MongoConfig) GetAllEventService {
	return &getalleventService{con: con}
}

func (s *getalleventService) GetAllEvent(ctx context.Context, filterReq models.EventFilterRequest) ([]models.Event, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filters := bson.M{}

	// Category
	if filterReq.Category != "" {
		filters["category"] = filterReq.Category
	}

	// Date (full day match)
	if filterReq.Date != "" {
		if date, err := time.Parse("2006-01-02", filterReq.Date); err == nil {
			startOfDay := date
			endOfDay := date.Add(24 * time.Hour)
			filters["startDateTime"] = bson.M{
				"$gte": startOfDay,
				"$lt":  endOfDay,
			}
		}
	}

	// Pagination
	page := filterReq.Page
	if page < 1 {
		page = 1
	}
	limit := filterReq.Limit
	if limit <= 0 {
		limit = 10
	}

	// Sorting
	sortField := "startDateTime"
	if filterReq.SortBy != "" {
		sortField = filterReq.SortBy
	}
	order := 1
	if strings.ToLower(filterReq.Order) == "desc" {
		order = -1
	}

	findOptions := options.Find().
		SetSkip(int64((page - 1) * limit)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: sortField, Value: order}}).
		SetProjection(bson.M{
			"title":          1,
			"category":       1,
			"startDateTime":  1,
			"bannerImageUrl": 1,
			"venueName":      1,
		})

	// Query DB
	collection := s.con.EventCol
	cursor, err := collection.Find(timeoutCtx, filters, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(timeoutCtx)

	// Decode into slice
	var events []models.Event
	if err := cursor.All(timeoutCtx, &events); err != nil {
		return nil, err
	}

	return events, nil
}
