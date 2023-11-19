package boarding_pass

import (
	"context"
	"github.com/nikolaevv/airtraffic/internal/model"
)

type CreateAdaptor interface {
	Create(ctx context.Context, flightID, seatID int) (model.BoardingPass, error)
}

type Create struct {
	repo CreateAdaptor
}

func NewCreate(repo CreateAdaptor) Create {
	return Create{
		repo: repo,
	}
}

func (act Create) Do(ctx context.Context, flightID, seatID int) (model.BoardingPass, error) {
	return act.repo.Create(ctx, flightID, seatID)
}
