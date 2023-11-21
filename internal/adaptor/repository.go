package adaptor

import (
	"context"

	"github.com/nikolaevv/airtraffic/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) GetBooking(ctx context.Context, id int) (model.Booking, error) {
	booking := model.Booking{}

	query := `
		select r.id, r.created_at, r.total_amount, t.id, t.passenger_id
		from bookings r
		left join tickets t on r.id = t.booking_id
		where r.id = $1
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return booking, errors.Wrap(err, "query booking")
	}

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
		booking.Tickets[i].Passenger, err = r.getPassenger(ctx, ticket.Passenger.ID)
		if err != nil {
			return booking, errors.Wrap(err, "get passenger")
		}

		booking.Tickets[i].Flights, err = r.getFlights(ctx, ticket.ID)
		if err != nil {
			return booking, errors.Wrap(err, "get flights")
		}
	}

	return booking, nil
}

func (r *Repository) getPassenger(ctx context.Context, id int) (model.Passenger, error) {
	var passenger model.Passenger

	var query = `
		select id, fullname, email
		from passengers
		where id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(&passenger.ID, &passenger.Name, &passenger.Email)
	if err != nil {
		return passenger, errors.Wrap(err, "query row passenger")
	}

	return passenger, nil
}

func (r *Repository) getFlights(ctx context.Context, ticketID int) ([]model.TicketFlight, error) {
	var ticketFlights []model.TicketFlight

	var query = `
		select t.id, t.amount, f.id, f.scheduled_departure, f.scheduled_arrival, f.departure_airport_id, f.arrival_airport_id, f.status, f.aircraft_id, f.actual_departure, f.actual_arrival
		from ticket_flights t
		left join flights f on t.flight_id = f.id
		where t.ticket_id = $1
	`

	rows, err := r.db.Query(ctx, query, ticketID)
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

func (r *Repository) CreateBooking(
	ctx context.Context,
	totalAmount float64,
	tickets []model.Ticket,
) (model.Booking, error) {
	booking := model.Booking{
		TotalAmount: totalAmount,
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return booking, errors.Wrap(err, "begin transaction")
	}

	err = tx.
		QueryRow(ctx, "insert into bookings (total_amount) values ($1) returning id, created_at", totalAmount).
		Scan(&booking.ID, &booking.CreatedAt)

	if err != nil {
		_ = tx.Rollback(ctx)
		return booking, errors.Wrap(err, "scan insert booking")
	}

	for _, ticket := range tickets {
		ticket, err := r.createTicket(ctx, tx, booking.ID, ticket.Passenger.ID)
		if err != nil {
			_ = tx.Rollback(ctx)
			return booking, errors.Wrap(err, "create ticket")
		}

		for i, flight := range ticket.Flights {
			ticketFlight, err := r.createTicketFlight(ctx, tx, ticket.ID, flight.Flight.ID, flight.Amount)
			if err != nil {
				_ = tx.Rollback(ctx)
				return booking, errors.Wrap(err, "create ticket flight")
			}

			ticket.Flights[i] = ticketFlight
		}

		booking.Tickets = append(booking.Tickets, ticket)
	}

	err = tx.Commit(ctx)
	return booking, err
}

func (r *Repository) createTicket(
	ctx context.Context,
	tx pgx.Tx,
	bookingID,
	passengerID int,
) (model.Ticket, error) {
	var ticket model.Ticket

	var query = `
		insert into tickets (booking_id, passenger_id)
		values ($1, $2)
		returning id
	`

	err := tx.QueryRow(ctx, query, bookingID, passengerID).Scan(&ticket.ID)
	if err != nil {
		return ticket, errors.Wrap(err, "scan insert ticket")
	}

	return ticket, nil
}

func (r *Repository) createTicketFlight(
	ctx context.Context,
	tx pgx.Tx,
	ticketID,
	flightID int,
	amount float64,
) (model.TicketFlight, error) {
	ticketFlight := model.TicketFlight{
		Amount: amount,
	}

	var query = `
		insert into ticket_flights (ticket_id, flight_id, amount)
		values ($1, $2, $3)
	    returning id
	`

	err := tx.QueryRow(ctx, query, ticketID, flightID, amount).Scan(&ticketFlight.ID)

	if err != nil {
		return ticketFlight, errors.Wrap(err, "scan insert ticket flight")
	}

	return ticketFlight, nil
}

func (r *Repository) CreateBoardingPass(ctx context.Context, flightID, seatID int) (model.BoardingPass, error) {
	boardingPass := model.BoardingPass{
		Seat: model.Seat{
			ID: seatID,
		},
	}

	var query = `
		insert into boarding_passes (flight_id, seat_id)
		values ($1, $2)
		returning id
	`

	err := r.db.QueryRow(ctx, query, flightID, seatID).Scan(&boardingPass.ID)

	if err != nil {
		return boardingPass, errors.Wrap(err, "scan insert boarding pass")
	}

	return boardingPass, nil
}

func (r *Repository) GetFlights(ctx context.Context) ([]model.Flight, error) {
	flights := make([]model.Flight, 0)

	var query = `
		select id, scheduled_departure, scheduled_arrival, status, aircraft_id, actual_departure, actual_arrival
		from flights
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "query select flights")
	}

	for rows.Next() {
		var flight model.Flight

		err := rows.Scan(
			&flight.ID,
			&flight.ScheduledDeparture,
			&flight.ScheduledArrival,
			&flight.Status,
			&flight.AircraftID,
			&flight.ActualDeparture,
			&flight.ActualArrival,
		)

		if err != nil {
			return nil, errors.Wrap(err, "scan select flights")
		}

		flights = append(flights, flight)
	}

	return flights, nil
}
