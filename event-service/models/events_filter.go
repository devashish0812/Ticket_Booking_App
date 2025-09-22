package models

// EventFilterRequest represents query parameters for listing events
type EventFilterRequest struct {
	Category string   `form:"category"` // ?category=Music&category=Comedy
	Tags     []string `form:"tag"`      // ?tag=EDM&tag=Bollywood
	Date     string   `form:"date"`     // ?date=2025-09-17 (ISO string or YYYY-MM-DD)
	SortBy   string   `form:"sortBy"`   // ?sortBy=startDateTime
	Order    string   `form:"order"`    // ?order=asc | desc
	Page     int      `form:"page,default=1"`
	Limit    int      `form:"limit,default=10"`
}
