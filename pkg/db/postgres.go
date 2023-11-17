package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
}

func Init(cfg Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.GetString("db.user"),
		cfg.GetString("db.pass"),
		cfg.GetString("db.host"),
		cfg.GetString("db.port"),
		cfg.GetString("db.name"),
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err = conn.Ping(context.Background()); err != nil {
		return nil, errors.WithStack(err)
	} else {
		log.Println("DB connection established")
	}

	return conn, nil
}
