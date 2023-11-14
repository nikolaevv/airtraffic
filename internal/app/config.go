package app

import "github.com/nikolaevv/airtraffic/internal/adaptor"

func New(configPath string) (*App, error) {
	cont, err := adaptor.NewContainer(configPath)
	if err != nil {
		return nil, err
	}

	return &App{
		cont: cont,
	}, nil
}
