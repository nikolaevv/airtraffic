package flights

import (
	"context"

	"github.com/nikolaevv/airtraffic/internal/model"
)

type GetListAdaptor interface {
	GetFlights(ctx context.Context) ([]model.Flight, error)
}

type GetList struct {
	repo GetListAdaptor
}

func NewGetList(repo GetListAdaptor) GetList {
	return GetList{
		repo: repo,
	}
}

func (act GetList) Do(ctx context.Context) ([]model.Flight, error) {
	return act.repo.GetFlights(ctx)
}
