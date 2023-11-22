package model

type Passenger struct {
	ID    int    `json:"id"`
	Name  string `json:"name" db:"fullname"`
	Email string `json:"email"`
}
