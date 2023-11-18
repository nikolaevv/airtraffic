package adaptor

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/nikolaevv/airtraffic/internal/model"
	"github.com/pkg/errors"
)

func NewBookingRepository(db *pgx.Conn) *BookingRepository {
	return &BookingRepository{db: db}
}

type BookingRepository struct {
	db *pgx.Conn
}

func (b *BookingRepository) Get(ctx context.Context, id int) (model.Booking, error) {
	booking := model.Booking{}

	query := `
		select b.id, b.created_at, b.total_amount, t.id, t.passenger_id
		from bookings b
		left join tickets t on b.id = t.booking_id
		where b.id = $1
	`

	rows, err := b.db.Query(ctx, query, id)
	if err != nil {
		return booking, errors.Wrap(err, "query booking")
	}
	//defer rows.Close()

	for rows.Next() {
		var ticket model.Ticket

		err := rows.Scan(&booking.ID, &booking.CreatedAt, &booking.TotalAmount, &ticket.ID, &ticket.Passenger.ID)
		if err != nil {
			return booking, errors.Wrap(err, "scan booking")
		}

		booking.Tickets = append(booking.Tickets, ticket)
	}

	rows.Close()
	for i, ticket := range booking.Tickets {
		booking.Tickets[i].Passenger, err = b.getPassenger(ctx, ticket.Passenger.ID)
		if err != nil {
			return booking, errors.Wrap(err, "get passenger")
		}

		booking.Tickets[i].Flights, err = b.getFlights(ctx, ticket.ID)
		if err != nil {
			return booking, errors.Wrap(err, "get flights")
		}
	}

	return booking, nil
}

func (b *BookingRepository) getPassenger(ctx context.Context, id int) (model.Passenger, error) {
	var passenger model.Passenger

	var query = `
		select id, fullname, email
		from passengers
		where id = $1
	`

	err := b.db.QueryRow(ctx, query, id).Scan(&passenger.ID, &passenger.Name, &passenger.Email)
	if err != nil {
		return passenger, errors.Wrap(err, "query row passenger")
	}

	return passenger, nil
}

func (b *BookingRepository) getFlights(ctx context.Context, ticketID int) ([]model.TicketFlight, error) {
	var ticketFlights []model.TicketFlight

	var query = `
		select t.id, t.amount, f.id, f.scheduled_departure, f.scheduled_arrival, f.departure_airport_id, f.arrival_airport_id, f.status, f.aircraft_id, f.actual_departure, f.actual_arrival
		from ticket_flights t
		left join flights f on t.flight_id = f.id
		where t.ticket_id = $1
	`

	rows, err := b.db.Query(ctx, query, ticketID)
	if err != nil {
		return ticketFlights, errors.Wrap(err, "query flights")
	}
	defer rows.Close()

	for rows.Next() {
		var ticketFlight model.TicketFlight

		err := rows.Scan(
			&ticketFlight.ID,
			&ticketFlight.Amount,
			&ticketFlight.Flight.ID,
			&ticketFlight.Flight.ScheduledDeparture,
			&ticketFlight.Flight.ScheduledArrival,
			&ticketFlight.Flight.DepartureAirportID,
			&ticketFlight.Flight.ArrivalAirportID,
			&ticketFlight.Flight.Status,
			&ticketFlight.Flight.AircraftID,
			&ticketFlight.Flight.ActualDeparture,
			&ticketFlight.Flight.ActualArrival,
		)

		if err != nil {
			return ticketFlights, errors.Wrap(err, "scan flights")
		}

		ticketFlights = append(ticketFlights, ticketFlight)
	}

	return ticketFlights, nil

}
