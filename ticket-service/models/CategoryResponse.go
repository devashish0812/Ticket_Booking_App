package models

type CategoryResponse struct {
	CategoryName     string  `json:"category_name" bson:"name"`
	Price            float64 `json:"price" bson:"price"`
	Description      string  `json:"description" bson:"description,omitempty"` // Note: This will be empty as it's not in Seat model
	TotalTickets     int     `json:"total_tickets" bson:"capacity"`
	BookedTickets    int     `json:"booked_tickets" bson:"booked_tickets"`
	AvailableTickets int     `json:"available_tickets" bson:"available_tickets"`
}
