package models

import "time"

type Master struct {
	Event   Event    `json:"event" binding:"required"`
	Tickets []Ticket `json:"tickets" binding:"required"`
}

type Event struct {
	ID          string   `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string   `bson:"title" json:"title"`
	Description string   `bson:"description" json:"description"`
	Category    string   `bson:"category" json:"category"`
	Tags        []string `bson:"tags,omitempty" json:"tags,omitempty"`

	VenueName string `bson:"venueName" json:"venueName"`
	Address   string `bson:"address" json:"address"`
	City      string `bson:"city" json:"city"`
	State     string `bson:"state" json:"state"`
	Country   string `bson:"country" json:"country"`
	Pincode   string `bson:"pincode" json:"pincode"`

	StartDateTime time.Time `bson:"startDateTime" json:"startDateTime"`
	EndDateTime   time.Time `bson:"endDateTime" json:"endDateTime"`
	Timezone      string    `bson:"timezone,omitempty" json:"timezone,omitempty"`

	BannerImageUrl string `bson:"bannerImageUrl,omitempty" json:"bannerImageUrl,omitempty"`

	MaxTicketsPerUser int `bson:"maxTicketsPerUser,omitempty" json:"maxTicketsPerUser,omitempty"`

	AgeRestriction string `bson:"ageRestriction,omitempty" json:"ageRestriction,omitempty"`
	Language       string `bson:"language,omitempty" json:"language,omitempty"`

	RefundPolicy string `bson:"refundPolicy,omitempty" json:"refundPolicy,omitempty"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type Ticket struct {
	EventId   string    `json:"eventId,omitempty"`
	Type      string    `json:"type" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
