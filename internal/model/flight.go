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
	ID                 int          `json:"id"`
	ScheduledDeparture time.Time    `json:"scheduled_departure"`
	ScheduledArrival   time.Time    `json:"scheduled_arrival"`
	DepartureAirport   Airport      `json:"departure_airport"`
	ArrivalAirport     Airport      `json:"arrival_airport"`
	Status             FlightStatus `json:"status"`
	AircraftID         int          `json:"aircraft_id"`
	ActualDeparture    time.Time    `json:"actual_departure"`
	ActualArrival      time.Time    `json:"actual_arrival"`
}
