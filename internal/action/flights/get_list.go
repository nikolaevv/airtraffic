package flights

import (
	"context"
	"time"

	"github.com/nikolaevv/airtraffic/internal/model"
)

//go:generate mockgen -source=get_list.go -destination=mock/get_list_mock.go -package=mock

type GetListAdaptor interface {
	GetFlights(ctx context.Context, date time.Time) ([]model.Flight, error)
}

type GetList struct {
	repo GetListAdaptor
}

func NewGetList(repo GetListAdaptor) GetList {
	return GetList{
		repo: repo,
	}
}

func (act GetList) Do(ctx context.Context, date time.Time) ([]model.Flight, error) {
	return act.repo.GetFlights(ctx, date)
}
