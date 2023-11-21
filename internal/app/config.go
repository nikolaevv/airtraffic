package app

import (
	"github.com/nikolaevv/airtraffic/internal/adaptor"

	"github.com/pkg/errors"
)

func New(configPath string) (*Server, error) {
	cont, err := adaptor.NewContainer(configPath)
	if err != nil {
		return nil, errors.New("init container")
	}

	return &Server{
		cont: cont,
	}, nil
}
