package adaptor

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/nikolaevv/airtraffic/internal/model"
	"github.com/pkg/errors"
)

func NewBoardingPassRepository(db *pgx.Conn) *BoardingPassRepository {
	return &BoardingPassRepository{db: db}
}

type BoardingPassRepository struct {
	db *pgx.Conn
}

func (b *BoardingPassRepository) Create(ctx context.Context, flightID, seatID int) (model.BoardingPass, error) {
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

	err := b.db.QueryRow(ctx, query, flightID, seatID).Scan(&boardingPass.ID)

	if err != nil {
		return boardingPass, errors.Wrap(err, "query boarding pass")
	}

	return boardingPass, nil
}
