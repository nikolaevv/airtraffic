package app

import "github.com/nikolaevv/airtraffic/internal/adaptor"

type App struct {
	cont adaptor.Container
}

func (s *App) Start() error {
	return nil
}
