package model

type TicketFlight struct {
	ID     int     `json:"id"`
	Flight Flight  `json:"flight"`
	Amount float64 `json:"amount"`
}
