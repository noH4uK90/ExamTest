package repository

import (
	"database/sql"
	"errors"

	"Api/internal/domains/models"
	"Api/internal/middleware"

	"github.com/jmoiron/sqlx"
)

type TestTypeRepository interface {
	Get(tx *sqlx.Tx) (*[]models.TestType, error)
	GetById(tx *sqlx.Tx, Id int64) (*models.TestType, error)
}

type TestTypeService struct {
	db *sqlx.DB
}

func NewTestTypeService(db *sqlx.DB) *TestTypeService {
	return &TestTypeService{
		db: db,
	}
}

func (s *TestTypeService) Get(tx *sqlx.Tx) (*[]models.TestType, error) {
	var testTypes []models.TestType

	err := tx.Select(&testTypes, `SELECT * FROM test_type ORDER BY "type_id"`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, middleware.NotFound
	}

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &testTypes, nil
}

func (s *TestTypeService) GetById(tx *sqlx.Tx, Id int64) (*models.TestType, error) {
	var testType models.TestType

	err := tx.Get(&testType, `SELECT * FROM test_type WHERE "type_id"=$1`, Id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, middleware.NotFound
	}

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &testType, nil
}
