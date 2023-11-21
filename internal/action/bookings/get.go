package bookings

import (
	"context"

	"github.com/nikolaevv/airtraffic/internal/model"
)

type GetAdaptor interface {
	GetBooking(ctx context.Context, id int) (model.Booking, error)
}

type Get struct {
	repo GetAdaptor
}

func NewGet(repo GetAdaptor) Get {
	return Get{
		repo: repo,
	}
}

func (act Get) Do(ctx context.Context, id int) (model.Booking, error) {
	return act.repo.GetBooking(ctx, id)
}
