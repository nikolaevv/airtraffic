package bookings_test

import (
	"context"
	"testing"

	"github.com/nikolaevv/airtraffic/internal/action/bookings"
	"github.com/nikolaevv/airtraffic/internal/action/bookings/mock"
	"github.com/nikolaevv/airtraffic/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Run("error GetBooking", func(t *testing.T) {
		repo := mock.NewMockGetAdaptor(gomock.NewController(t))

		repo.
			EXPECT().
			GetBooking(gomock.Any(), gomock.Any()).
			Return(model.Booking{}, assert.AnError)

		act := bookings.NewGet(repo)
		booking, err := act.Do(context.Background(), 1)

		assert.Equal(t, model.Booking{}, booking)
		assert.ErrorIs(t, err, assert.AnError)
	})

	t.Run("success", func(t *testing.T) {
		repo := mock.NewMockGetAdaptor(gomock.NewController(t))

		expectedBooking := model.Booking{
			ID: 1,
		}

		repo.
			EXPECT().
			GetBooking(gomock.Any(), gomock.Any()).
			Return(expectedBooking, nil)

		act := bookings.NewGet(repo)
		booking, err := act.Do(context.Background(), 1)

		assert.Equal(t, expectedBooking, booking)
		assert.NoError(t, err)
	})
}
