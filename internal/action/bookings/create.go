package bookings

import (
	"context"
	"github.com/nikolaevv/airtraffic/internal/model"
	"github.com/pkg/errors"
)

const (
	singleTicketPrice = 100
)

type CreateAdaptor interface {
	CreateBooking(ctx context.Context, totalAmount float64) (model.Booking, error)
	CreateTicket(ctx context.Context, bookingID, passengerID int) (model.Ticket, error)
	CreateFlightTicket(ctx context.Context, ticketID, flightID int, amount float64) (model.TicketFlight, error)
}

type Create struct {
	repo CreateAdaptor
}

func NewCreate(repo CreateAdaptor) Create {
	return Create{
		repo: repo,
	}
}

func (act Create) Do(ctx context.Context, flightID int, passengers []model.Passenger) (model.Booking, error) {
	total := float64(len(passengers) * singleTicketPrice)

	booking, err := act.repo.CreateBooking(ctx, total)
	if err != nil {
		return booking, errors.Wrap(err, "create booking")
	}

	for _, passenger := range passengers {
		ticket, err := act.repo.CreateTicket(ctx, booking.ID, passenger.ID)
		if err != nil {
			return booking, errors.Wrap(err, "create ticket")
		}

		ticket.Passenger = passenger
		booking.Tickets = append(booking.Tickets, ticket)
	}

	for i, ticket := range booking.Tickets {
		flightTicket, err := act.repo.CreateFlightTicket(ctx, ticket.ID, flightID, singleTicketPrice)
		if err != nil {
			return booking, errors.Wrap(err, "create flight ticket")
		}

		booking.Tickets[i].Flights = append(booking.Tickets[i].Flights, flightTicket)
	}

	return booking, nil
}
