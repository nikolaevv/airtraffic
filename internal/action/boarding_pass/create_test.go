package boarding_pass_test

import (
	"context"
	"testing"

	"github.com/nikolaevv/airtraffic/internal/action/boarding_pass"
	"github.com/nikolaevv/airtraffic/internal/action/boarding_pass/mock"
	"github.com/nikolaevv/airtraffic/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	t.Run("error CreateBoardingPass", func(t *testing.T) {
		repo := mock.NewMockCreateAdaptor(gomock.NewController(t))

		repo.
			EXPECT().
			CreateBoardingPass(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(model.BoardingPass{}, assert.AnError)

		act := boarding_pass.NewCreate(repo)
		boardingPass, err := act.Do(ctx, 1, 1)

		assert.Equal(t, model.BoardingPass{}, boardingPass)
		assert.ErrorIs(t, err, assert.AnError)
	})

	t.Run("success", func(t *testing.T) {
		repo := mock.NewMockCreateAdaptor(gomock.NewController(t))

		expectedBoardingPass := model.BoardingPass{
			Seat: model.Seat{
				ID: 1,
			},
		}

		repo.
			EXPECT().
			CreateBoardingPass(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(expectedBoardingPass, nil)

		act := boarding_pass.NewCreate(repo)
		boardingPass, err := act.Do(ctx, 1, 1)

		assert.Equal(t, expectedBoardingPass, boardingPass)
		assert.NoError(t, err)
	})
}
