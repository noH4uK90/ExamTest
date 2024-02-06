package app

import (
	"Api/internal/config"
	"context"
	"github.com/go-chi/chi/v5"
	"log"
)

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

func (a *App) initDeps(ctx *context.Context) error {
	return nil
}

func (a *App) initEnv(_ *context.Context) error {
	err := config.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loadl config: %s", err.Error())
	}
	return nil
}

func (a *App) initServiceProvider(_ *context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func initRouter(a *App) (*chi.Mux, error) {
	r := chi.NewRouter()
	//postgres := a.serviceProvider.Postgres()
	return r, nil
}
