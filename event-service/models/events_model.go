package models

import "time"

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

type TicketSummary struct {
	Type     string  `bson:"type" json:"type"`
	Price    float64 `bson:"price" json:"price"`
	Currency string  `bson:"currency" json:"currency"`
}
