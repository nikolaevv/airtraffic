package model

import "time"

type FlightStatus string

const (
	FlightStatusScheduled FlightStatus = "SCHEDULED"
	FlightStatusOnTime    FlightStatus = "ONTIME"
	FlightStatusDelayed   FlightStatus = "DELAYED"
	FlightStatusCancelled FlightStatus = "CANCELLED"
	FlightStatusDiverted  FlightStatus = "DEPARTED"
	FlightStatusUnknown   FlightStatus = "ARRIVED"
)

type Flight struct {
	ID                 int       `json:"id"`
	ScheduledDeparture time.Time `json:"scheduled_departure"`
	ScheduledArrival   time.Time `json:"scheduled_arrival"`
	//DepartureAirport   string    `json:"departure_airport"`
	//ArrivalAirport     string    `json:"arrival_airport"`
	Status          FlightStatus `json:"status"`
	AircraftCode    string       `json:"aircraft_code"`
	ActualDeparture time.Time    `json:"actual_departure"`
	ActualArrival   time.Time    `json:"actual_arrival"`
}
