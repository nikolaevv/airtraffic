package model

import "time"

type Booking struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	TotalAmount float64   `json:"total_amount"`
	Tickets     []Ticket  `json:"tickets"`
}
