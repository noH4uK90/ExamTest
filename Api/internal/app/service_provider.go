package app

import (
	"Api/internal/config"
	"Api/internal/domains/database"
	"Api/internal/repository"
	"github.com/jmoiron/sqlx"
	"log"
)

type ServiceProvider struct {
	postgres         *sqlx.DB
	answerRepository repository.AnswerRepository
}

func newServiceProvider() *ServiceProvider { return &ServiceProvider{} }

func (s *ServiceProvider) Postgres() *sqlx.DB {
	if s.postgres == nil {
		cfg, err := config.NewPostgresConfig()

		if err != nil {
			log.Fatalf("Failed to get postgres config: %s", err.Error())
		}

		s.postgres, err = database.NewPostgresConnection(cfg)

		if err != nil {
			log.Fatalf("Failed connect to postgres: %s", err.Error())
		}
	}

	return s.postgres
}

func (s *ServiceProvider) AnswerService() repository.AnswerRepository {
	if s.answerRepository == nil {
		s.answerRepository = repository.NewAnswerService(s.postgres)
	}

	return s.answerRepository
}
