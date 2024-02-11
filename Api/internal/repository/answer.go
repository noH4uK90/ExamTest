package repository

import (
	"database/sql"
	"errors"

	"Api/internal/domains/models"
	"Api/internal/middleware"

	"github.com/jmoiron/sqlx"
)

type AnswerRepository interface {
	GetById(tx *sqlx.Tx, ID int64) (*models.Answer, error)
	Get(tx *sqlx.Tx) (*[]models.Answer, error)
	Create(tx *sqlx.Tx, answer models.Answer) (*int64, error)
	Update(tx *sqlx.Tx, ID int64, answer models.Answer) error
	Delete(tx *sqlx.Tx, ID int64) error
}

type AnswerService struct {
	db *sqlx.DB
}

func NewAnswerService(db *sqlx.DB) *AnswerService {
	return &AnswerService{
		db: db,
	}
}

func (s *AnswerService) GetById(tx *sqlx.Tx, ID int64) (*models.Answer, error) {
	var answer models.Answer

	err := tx.Get(&answer, `SELECT * FROM answer WHERE "answer_id" = $1`, ID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, middleware.NotFound
	}

	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &answer, nil
}

func (s *AnswerService) Get(tx *sqlx.Tx) (*[]models.Answer, error) {
	var answers []models.Answer

	err := tx.Select(&answers, `SELECT * FROM answer ORDER BY "answer_id"`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, middleware.NotFound
	}

	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &answers, nil
}

func (s *AnswerService) Create(tx *sqlx.Tx, answer models.Answer) (*int64, error) {
	var id int64
	var isExists bool

	err := tx.Get(&isExists, `SELECT EXISTS(SELECT * FROM answer WHERE "text"=$1)`, answer.Text)
	if isExists == true {
		return nil, middleware.IsExist
	}
	if err != nil {
		return nil, err
	}

	err = tx.QueryRowx(`INSERT INTO answer("text", "is_right") VALUES($1, $2) RETURNING "answer_id"`, answer.Text, answer.IsRight).Scan(&id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (s *AnswerService) Update(tx *sqlx.Tx, ID int64, answer models.Answer) error {

	_, err := tx.Queryx(`UPDATE answer SET "text"=$1, "is_right"=$2 WHERE "answer_id"=$3`, answer.Text, answer.IsRight, ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *AnswerService) Delete(tx *sqlx.Tx, ID int64) error {
	res := tx.MustExec(`DELETE FROM answer WHERE "answer_id" = $1`, ID)
	_, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
