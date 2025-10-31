package services

import (
	"context"
	"strings"
	"time"

	"github.com/devashish0812/event-service/config"
	"github.com/devashish0812/event-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetAllEventForOrgService interface {
	GetAllEventForOrg(ctx context.Context, filterReq models.EventFilterRequest) ([]models.Event, error)
}

type getalleventForOrgService struct {
	con *config.MongoConfig
}

func NewGetAllEventForOrgService(con *config.MongoConfig) GetAllEventForOrgService {
	return &getalleventForOrgService{con: con}
}

func (s *getalleventForOrgService) GetAllEventForOrg(ctx context.Context, filterReq models.EventFilterRequest) ([]models.Event, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filters := bson.M{}

	// Category
	if filterReq.Type != "" {
		if strings.EqualFold(filterReq.Type, "upcoming") {
			filters["startDateTime"] = bson.M{"$gt": time.Now()}
		} else if strings.EqualFold(filterReq.Type, "ongoing") {
			filters["startDateTime"] = bson.M{"$lte": time.Now()}
			filters["endDateTime"] = bson.M{"$gte": time.Now()}
		} else if strings.EqualFold(filterReq.Type, "past") {
			filters["endDateTime"] = bson.M{"$lt": time.Now()}
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

	findOptions := options.Find().
		SetSkip(int64((page - 1) * limit)).
		SetLimit(int64(limit)).
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
