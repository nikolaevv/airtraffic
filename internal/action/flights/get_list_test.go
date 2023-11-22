package flights_test

import (
	"context"
	"testing"

	"github.com/nikolaevv/airtraffic/internal/action/flights"
	"github.com/nikolaevv/airtraffic/internal/action/flights/mock"
	"github.com/nikolaevv/airtraffic/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	ctx := context.Background()

	t.Run("error GetFlights", func(t *testing.T) {
		repo := mock.NewMockGetListAdaptor(gomock.NewController(t))

		repo.
			EXPECT().
			GetFlights(gomock.Any()).
			Return(nil, assert.AnError)

		act := flights.NewGetList(repo)
		_, err := act.Do(ctx)

		assert.ErrorIs(t, err, assert.AnError)
	})

	t.Run("success", func(t *testing.T) {
		repo := mock.NewMockGetListAdaptor(gomock.NewController(t))

		flightsList := []model.Flight{
			{
				ID: 1,
			},
		}

		repo.
			EXPECT().
			GetFlights(gomock.Any()).
			Return(flightsList, nil)

		act := flights.NewGetList(repo)
		res, err := act.Do(ctx)

		assert.NoError(t, err)
		assert.Equal(t, flightsList, res)
	})
}
