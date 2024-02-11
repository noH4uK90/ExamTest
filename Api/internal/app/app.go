package app

import (
	"context"
	"log"
	"net/http"

	"Api/internal/config"
	"Api/internal/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type App struct {
	serviceProvider *ServiceProvider
}

func NewApp(ctx *context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	cfg, err := config.NewApiConfiguration()
	if err != nil {
		return err
	}

	r, err := initRouter(a)
	if err != nil {
		return err
	}

	err = http.ListenAndServe(cfg.Address(), r)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps(ctx *context.Context) error {
	inits := []func(*context.Context) error{
		a.initEnv,
		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

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
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	postgres := a.serviceProvider.Postgres()

	r.Route("/api", func(r chi.Router) {
		controller.NewAnswerController().Init(r, a.serviceProvider.AnswerService(), postgres)
		controller.NewTestController().Init(r, a.serviceProvider.TestService(), postgres)
		controller.NewTestTypeController().Init(r, a.serviceProvider.TestTypeService(), postgres)
	})

	return r, nil
}
