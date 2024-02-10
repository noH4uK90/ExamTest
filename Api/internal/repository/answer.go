package repository

import (
	"Api/internal/domains/models"
	"Api/internal/middleware"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type AnswerRepository interface {
	GetById(tx *sqlx.Tx, Id int64) (*models.Answer, error)
	Get(tx *sqlx.Tx) (*[]models.Answer, error)
	Create(tx *sqlx.Tx, answer models.Answer) (*int64, error)
	Update(tx *sqlx.Tx, Id int64, answer models.Answer) error
	Delete(tx *sqlx.Tx, Id int64) error
}

type AnswerService struct {
	db *sqlx.DB
}

func NewAnswerService(db *sqlx.DB) *AnswerService {
	return &AnswerService{
		db: db,
	}
}

func (r *AnswerService) GetById(tx *sqlx.Tx, Id int64) (*models.Answer, error) {
	var answer models.Answer

	err := tx.Get(&answer, `SELECT * FROM answer WHERE "answer_id" = $1`, Id)
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

func (r *AnswerService) Get(tx *sqlx.Tx) (*[]models.Answer, error) {
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

func (r *AnswerService) Create(tx *sqlx.Tx, answer models.Answer) (*int64, error) {
	var id int64
	var ans models.Answer

	err := tx.Get(&ans, `SELECT * FROM answer WHERE "text"=$1`, answer.Text)
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, middleware.IsExist
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

func (r *AnswerService) Update(tx *sqlx.Tx, Id int64, answer models.Answer) error {

	_, err := tx.Queryx(`UPDATE answer SET "text"=$1, "is_right"=$2 WHERE "answer_id"=$3`, answer.Text, answer.IsRight, Id)
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

func (r *AnswerService) Delete(tx *sqlx.Tx, Id int64) error {
	res := tx.MustExec(`DELETE FROM answer WHERE "answer_id" = $1`, Id)
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
