package models

type SectionResponse struct {
	SectionName      string `json:"section_name" bson:"section_name"`
	Capacity         int    `json:"capacity" bson:"capacity"`
	AvailableTickets int    `json:"available_tickets" bson:"available_tickets"`
}
