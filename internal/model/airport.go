package model

type Airport struct {
	ID        int     `json:"id"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	City      string  `json:"city"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Timezone  string  `json:"timezone"`
}
