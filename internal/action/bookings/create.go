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
	CreateBooking(ctx context.Context, totalAmount float64, tickets []model.Ticket) (model.Booking, error)
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

	tickets := make([]model.Ticket, 0, len(passengers))

	for _, passenger := range passengers {
		ticket := model.Ticket{
			Passenger: passenger,
			Flights: []model.TicketFlight{
				{
					Flight: model.Flight{
						ID: flightID,
					},
					Amount: singleTicketPrice,
				},
			},
		}

		tickets = append(tickets, ticket)
	}

	booking, err := act.repo.CreateBooking(ctx, total, tickets)
	if err != nil {
		return booking, errors.Wrap(err, "create booking")
	}

	return booking, nil
}
