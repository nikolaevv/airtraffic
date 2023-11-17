package app

import "github.com/nikolaevv/airtraffic/internal/adaptor"

func New(configPath string) (*Server, error) {
	cont, err := adaptor.NewContainer(configPath)
	if err != nil {
		return nil, err
	}

	return &Server{
		cont: cont,
	}, nil
}
