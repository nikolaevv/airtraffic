package adaptor

import (
	"context"
	"strconv"
	"strings"
	"time"

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

func (r *Repository) GetBooking(ctx context.Context, id int) (booking model.Booking, err error) {
	bookingQuery := `
		select id, created_at, total_amount
		from bookings
		where id = $1
	`

	err = r.db.QueryRow(ctx, bookingQuery, id).Scan(&booking.ID, &booking.CreatedAt, &booking.TotalAmount)
	if err != nil {
		return booking, errors.Wrap(err, "query row booking")
	}

	// Only ticket IDs are needed
	ticketIDs, err := r.getTicketIDs(ctx, id)
	if err != nil {
		return booking, errors.Wrap(err, "get ticket ids")
	}

	ticketIDsFilter := r.buildTicketIDsFilter(ticketIDs)

	passengers, err := r.getPassengers(ctx, ticketIDs, ticketIDsFilter)
	if err != nil {
		return booking, errors.Wrap(err, "get passengers")
	}

	ticketFlights, err := r.getTicketFlights(ctx, ticketIDs, ticketIDsFilter)
	if err != nil {
		return booking, errors.Wrap(err, "get ticket flights")
	}

	for _, ticketID := range ticketIDs {
		booking.Tickets = append(booking.Tickets, model.Ticket{
			ID:        ticketID,
			Passenger: passengers[ticketID],
			Flights:   ticketFlights[ticketID],
		})
	}

	return booking, nil
}

func (r *Repository) getTicketIDs(ctx context.Context, bookingID int) ([]int, error) {
	var ticketIDs []int

	rows, err := r.db.Query(ctx, `
		select id
		from tickets
		where booking_id = $1
	`, bookingID)

	if err != nil {
		return ticketIDs, errors.Wrap(err, "query ticket_ids")
	}

	for rows.Next() {
		var ticketID int

		err = rows.Scan(&ticketID)
		if err != nil {
			return ticketIDs, errors.Wrap(err, "scan ticket_ids")
		}

		ticketIDs = append(ticketIDs, ticketID)
	}

	return ticketIDs, nil
}

func (r *Repository) buildTicketIDsFilter(ticketIDs []int) string {
	var ticketIDsFilter []string

	for _, ticketID := range ticketIDs {
		ticketIDsFilter = append(ticketIDsFilter, strconv.Itoa(ticketID))
	}

	return strings.Join(ticketIDsFilter, ",")
}

func (r *Repository) getTicketFlights(
	ctx context.Context,
	ticketIDs []int,
	ticketIDsFilter string,
) (map[int][]model.TicketFlight, error) {
	flights := make(map[int][]model.TicketFlight, len(ticketIDs))

	for _, ticketID := range ticketIDs {
		flights[ticketID] = []model.TicketFlight{}
	}

	rows, err := r.db.Query(ctx, `
		select t.id, t.amount, f.id, f.scheduled_departure, f.scheduled_arrival, f.departure_airport_id, f.arrival_airport_id, f.status, f.aircraft_id, f.actual_departure, f.actual_arrival
		from ticket_flights t
		left join flights f on t.flight_id = f.id
		where t.ticket_id in ($1)
	`, ticketIDsFilter)

	if err != nil {
		return flights, errors.Wrap(err, "query flights")
	}

	for rows.Next() {
		var ticketFlight model.TicketFlight
		var flight model.Flight

		err = rows.Scan(
			&ticketFlight.ID,
			&ticketFlight.Amount,
			&flight.ID,
			&flight.ScheduledDeparture,
			&flight.ScheduledArrival,
			&flight.DepartureAirportID,
			&flight.ArrivalAirportID,
			&flight.Status,
			&flight.AircraftID,
			&flight.ActualDeparture,
			&flight.ActualArrival,
		)

		if err != nil {
			return flights, errors.Wrap(err, "scan flights")
		}

		flights[ticketFlight.ID] = append(flights[ticketFlight.ID], ticketFlight)
	}

	return flights, nil
}

func (r *Repository) getPassengers(
	ctx context.Context,
	ticketIDs []int,
	ticketIDsFilter string,
) (map[int]model.Passenger, error) {
	passengers := make(map[int]model.Passenger, len(ticketIDs))

	rows, err := r.db.Query(ctx, `
		select id, fullname, email
		from passengers
		where id in ($1)
	`, ticketIDsFilter)

	if err != nil {
		return passengers, errors.Wrap(err, "query passengers")
	}

	for rows.Next() {
		passenger, err := pgx.RowToStructByName[model.Passenger](rows)
		if err != nil {
			return passengers, errors.Wrap(err, "scan passengers")
		}

		passengers[passenger.ID] = passenger
	}

	return passengers, nil
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

func (r *Repository) GetFlights(ctx context.Context, date time.Time) ([]model.Flight, error) {
	flights := make([]model.Flight, 0)

	var query = `
		select id, scheduled_departure, scheduled_arrival, departure_airport_id, arrival_airport_id, status, aircraft_id, actual_departure, actual_arrival
		from flights
		where scheduled_departure::date = $1
	`

	rows, err := r.db.Query(ctx, query, date)
	if err != nil {
		return nil, errors.Wrap(err, "query select flights")
	}

	for rows.Next() {
		flight, err := pgx.RowToStructByName[model.Flight](rows)

		if err != nil {
			return nil, errors.Wrap(err, "scan select flights")
		}

		flights = append(flights, flight)
	}

	return flights, nil
}
