package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Seat struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	EventID     string             `json:"eventId" bson:"eventId"`
	Category    string             `json:"category" bson:"category"`     // "Gold", "Silver"
	SeatNumber  string             `json:"seatNumber" bson:"seatNumber"` // "G-1", "G-2", "S-1"
	Row         string             `json:"row" bson:"row"`               // "A", "B", "C" (optional)
	Price       float64            `json:"price" bson:"price"`
	Status      string             `json:"status" bson:"status"`                               // "available", "locked", "booked"
	LockedBy    string             `json:"lockedBy,omitempty" bson:"lockedBy,omitempty"`       // UserID
	LockedUntil time.Time          `json:"lockedUntil,omitempty" bson:"lockedUntil,omitempty"` // Lock expiry
	BookedBy    string             `json:"bookedBy,omitempty" bson:"bookedBy,omitempty"`       // UserID
	BookedAt    time.Time          `json:"bookedAt,omitempty" bson:"bookedAt,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type Ticket struct {
	EventId   string    `json:"eventId,omitempty"`
	Type      string    `json:"type" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
