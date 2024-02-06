package answer

import (
	"Api/internal/domains/models"
	"Api/internal/middleware"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetById(tx *sqlx.Tx, Id int64) (*models.Answer, error)
	Get(tx *sqlx.Tx) (*[]models.Answer, error)
}

type AnswerRepository struct {
	db *sqlx.DB
}

func NewAnswerRepository(db *sqlx.DB) *AnswerRepository {
	return &AnswerRepository{
		db: db,
	}
}

func (r *AnswerRepository) GetById(tx *sqlx.Tx, Id int64) (*models.Answer, error) {
	var answer models.Answer

	err := tx.Get(answer, `SELECT * FROM answer WHERE "answer_id" = $1`, Id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, middleware.NotFound
	}

	if err != nil {
		return nil, err
	}

	return &answer, nil
}

func (r *AnswerRepository) Get(tx *sqlx.Tx) (*[]models.Answer, error) {
	var answers []models.Answer

	err := tx.Get(answers, `SELECT * FROM answer ORDER BY "answer_id"`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, middleware.NotFound
	}

	if err != nil {
		return nil, err
	}

	return &answers, nil
}
