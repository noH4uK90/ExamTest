package app

import (
	"log"

	"Api/internal/config"
	"Api/internal/domains/database"
	"Api/internal/repository"

	"github.com/jmoiron/sqlx"
)

type ServiceProvider struct {
	postgres           *sqlx.DB
	answerRepository   repository.AnswerRepository
	testTypeRepository repository.TestTypeRepository
	testRepository     repository.TestRepository
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

func (s *ServiceProvider) TestTypeService() repository.TestTypeRepository {
	if s.testTypeRepository == nil {
		s.testTypeRepository = repository.NewTestTypeService(s.postgres)
	}

	return s.testTypeRepository
}

func (s *ServiceProvider) TestService() repository.TestRepository {
	if s.testRepository == nil {
		s.testRepository = repository.NewTestService(s.postgres)
	}

	return s.testRepository
}
