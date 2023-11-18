package adaptor

import (
	"context"

	"github.com/nikolaevv/airtraffic/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func NewFlightRepository(db *pgx.Conn) *FlightRepository {
	return &FlightRepository{db: db}
}

type FlightRepository struct {
	db *pgx.Conn
}

func (f *FlightRepository) GetList(ctx context.Context) ([]model.Flight, error) {
	flights := make([]model.Flight, 0)

	var query = "select id, scheduled_departure, scheduled_arrival, status, aircraft_id, actual_departure, actual_arrival from flights"
	rows, err := f.db.Query(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "query flight list")
	}

	for rows.Next() {
		var flight model.Flight
		if err := rows.Scan(&flight.ID, &flight.ScheduledDeparture, &flight.ScheduledArrival, &flight.Status, &flight.AircraftID, &flight.ActualDeparture, &flight.ActualArrival); err != nil {
			return nil, errors.Wrap(err, "scan flight list")
		}

		flights = append(flights, flight)
	}

	return flights, nil
}
