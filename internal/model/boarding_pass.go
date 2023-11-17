package model

type BoardingPass struct {
	ID   int  `json:"id"`
	Seat Seat `json:"seat"`
}
