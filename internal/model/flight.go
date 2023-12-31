package model

import (
	"database/sql"
	"time"
)

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
	ScheduledDeparture time.Time    `json:"scheduled_departure" db:"scheduled_departure"`
	ScheduledArrival   time.Time    `json:"scheduled_arrival" db:"scheduled_arrival"`
	DepartureAirportID int          `json:"departure_airport_id" db:"departure_airport_id"`
	ArrivalAirportID   int          `json:"arrival_airport_id" db:"arrival_airport_id"`
	Status             FlightStatus `json:"status"`
	AircraftID         int          `json:"aircraft_id" db:"aircraft_id"`
	ActualDeparture    sql.NullTime `json:"actual_departure" db:"actual_departure"`
	ActualArrival      sql.NullTime `json:"actual_arrival" db:"actual_arrival"`
}
