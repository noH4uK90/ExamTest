package repository

import (
	"Api/internal/domains/models"
	"Api/internal/middleware"
	"database/sql"
	"reflect"

	"errors"
	"github.com/jmoiron/sqlx"
)

type TestRepository interface {
	Get(tx *sqlx.Tx) ([]uint8, error)
	GetById(tx *sqlx.Tx, ID int64) ([]uint8, error)
	Create(tx *sqlx.Tx, test models.Test) (*int64, error)
}

type TestService struct {
	db *sqlx.DB
}

func NewTestService(db *sqlx.DB) *TestService {
	return &TestService{
		db: db,
	}
}

func (s *TestService) Get(tx *sqlx.Tx) ([]uint8, error) {
	var rows uint8
	q := `
	SELECT json_agg(
    json_build_object(
        'testId', test.test_id,
        'name', name,
        'questions', array_to_json(array(
            SELECT json_build_object(
                'questionId', question.question_id,
                'text', question.text,
                'scoreId', question.score_id,
                'answers', array_to_json(array(
                    SELECT json_build_object(
                        'answerId', answer_id,
                        'text', answer.text,
                        'isRight', is_right
                    )
                    FROM answer
                    WHERE answer.answer_id IN (
                        SELECT answer_id
                        FROM question_answer
                        WHERE question_answer.question_id=question.question_id
                    )
                ))
            )
            FROM question
            WHERE question.question_id IN (
                SELECT question_id
                FROM test_question
                WHERE test_question.test_id = test.test_id)
            ))
    ))
	FROM test;`

	slice := reflect.New(reflect.SliceOf(reflect.TypeOf(rows)))
	err := tx.Get(slice.Interface(), q)
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

	return slice.Elem().Bytes(), nil
}

func (s *TestService) GetById(tx *sqlx.Tx, ID int64) ([]uint8, error) {
	var rows uint8
	q := `
	SELECT json_build_object(
               'testId', test.test_id,
               'name', name,
               'questions', array_to_json(array(
                SELECT json_build_object(
                               'questionId', question.question_id,
                               'text', question.text,
                               'scoreId', question.score_id,
                               'answers', array_to_json(array(
                                SELECT json_build_object(
                                               'answerId', answer_id,
                                               'text', answer.text,
                                               'isRight', is_right
                                       )
                                FROM answer
                                WHERE answer.answer_id IN (SELECT answer_id
                                                           FROM question_answer
                                                           WHERE question_answer.question_id = question.question_id)
                                                        ))
                       )
                FROM question
                WHERE question.question_id IN (SELECT question_id
                                               FROM test_question
                                               WHERE test_question.test_id = test.test_id)
                                          ))
       )
	FROM test
	WHERE test_id = $1`
	slice := reflect.New(reflect.SliceOf(reflect.TypeOf(rows)))

	err := tx.Get(slice.Interface(), q, ID)
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

	return slice.Elem().Bytes(), nil
}

func (s *TestService) Create(tx *sqlx.Tx, test models.Test) (*int64, error) {
	var id int64
	var isExists bool

	err := tx.Get(&isExists, `SELECT EXISTS(SELECT * FROM test WHERE "name"=$1)`, test.Name)
	if isExists == true {
		return nil, middleware.IsExist
	}
	if err != nil {
		return nil, err
	}

	err = tx.QueryRowx(`INSERT INTO test(name) VALUES ($1)`, test.Name).Scan(&id)
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
