package model

type Aircraft struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Seats []Seat `json:"seats"`
}
