package adaptor

import (
	"log"

	"github.com/nikolaevv/airtraffic/pkg/config"
	"github.com/nikolaevv/airtraffic/pkg/db"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Container struct {
	cfg *viper.Viper
	db  *pgx.Conn
}

func NewContainer(configPath string) (*Container, error) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "init config")
	}

	database, err := db.Init(cfg)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "init db")
	}

	return &Container{
		cfg: cfg,
		db:  database,
	}, nil
}

func (c Container) GetConfig() *viper.Viper {
	return c.cfg
}

func (c Container) GetRepository() *Repository {
	return NewRepository(c.db)
}
