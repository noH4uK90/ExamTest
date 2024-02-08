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
	Create(tx *sqlx.Tx, answer models.Answer) (*int64, error)
	Update(tx *sqlx.Tx, Id int64, answer models.Answer) error
	Delete(tx *sqlx.Tx, Id int64) error
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

	err := tx.Get(&answer, `SELECT * FROM answer WHERE "answer_id" = $1`, Id)
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

	err := tx.Select(&answers, `SELECT * FROM answer ORDER BY "answer_id"`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, middleware.NotFound
	}

	if err != nil {
		return nil, err
	}

	return &answers, nil
}

func (r *AnswerRepository) Create(tx *sqlx.Tx, answer models.Answer) (*int64, error) {
	var id int64

	err := tx.QueryRowx(`INSERT INTO answer("text", "is_right") VALUES($1, $2) RETURNING "answer_id"`, answer.Text, answer.IsRight).Scan(&id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &id, nil
}

func (r *AnswerRepository) Update(tx *sqlx.Tx, Id int64, answer models.Answer) error {
	_, err := tx.Queryx(`UPDATE answer SET "text"=$1, "is_right"=$2 WHERE "answer_id"=$3`, Id, answer.Text, answer.IsRight)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *AnswerRepository) Delete(tx *sqlx.Tx, Id int64) error {
	res := tx.MustExec(`DELETE FROM answer WHERE "answer_id" = $1`, Id)
	_, err := res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
