package model

type Ticket struct {
	ID        int       `json:"id"`
	Passenger Passenger `json:"passenger"`
	Flights   []Flight  `json:"flights"`
}
