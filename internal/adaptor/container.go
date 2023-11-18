package adaptor

import (
	"github.com/jackc/pgx/v5"
	"github.com/nikolaevv/airtraffic/pkg/config"
	"github.com/nikolaevv/airtraffic/pkg/db"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
)

type Container struct {
	cfg *viper.Viper
	db  *pgx.Conn
}

func NewContainer(configPath string) (*Container, error) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Println(err)
		return nil, errors.WithStack(err)
	}

	db, err := db.Init(cfg)
	if err != nil {
		log.Println(err)
		return nil, errors.WithStack(err)
	}

	return &Container{
		cfg: cfg,
		db:  db,
	}, nil
}

func (c Container) GetConfig() *viper.Viper {
	return c.cfg
}

func (c Container) GetDatabase() *pgx.Conn {
	return c.db
}

func (c Container) GetFlightRepository() *FlightRepository {
	return NewFlightRepository(c.db)
}

func (c Container) GetBookingRepository() *BookingRepository {
	return NewBookingRepository(c.db)
}
