package app

import "context"

type App struct {
	serviceProvider *ServiceProvider
}

func NewApp(ctx *context.Context) (*App, error) {
	a := &App{}

	return a, nil
}

func (a *App) Run() error {
	return nil
}
