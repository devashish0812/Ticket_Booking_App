package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OutboxEvent represents a record in the outbox collection.
type OutboxEvent struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EventID   string             `bson:"eventId" json:"eventId"`
	EventType string             `bson:"eventType" json:"eventType"`
	Payload   interface{}        `bson:"payload" json:"payload"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	Processed bool               `bson:"processed" json:"processed"`
}
