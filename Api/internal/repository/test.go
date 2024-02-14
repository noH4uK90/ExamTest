package repository

import (
	"Api/internal/domains/models"
	"Api/internal/middleware"
	"database/sql"
	"github.com/lib/pq"
	"reflect"

	"errors"
	"github.com/jmoiron/sqlx"
)

type TestRepository interface {
	Get(tx *sqlx.Tx) ([]uint8, error)
	GetById(tx *sqlx.Tx, ID int64) ([]uint8, error)
	Create(tx *sqlx.Tx, test models.TestRequest) (*int, error)
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
        'id', test.test_id,
        'name', name,
        'questions', array_to_json(array(
            SELECT json_build_object(
                'id', question.question_id,
                'text', question.text,
                'scoreId', question.score_id,
                'answers', array_to_json(array(
                    SELECT json_build_object(
                        'id', answer_id,
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
	q := `
	SELECT json_build_object(
               'id', test.test_id,
               'name', name,
               'questions', array_to_json(array(
                SELECT json_build_object(
                               'id', question.question_id,
                               'text', question.text,
                               'scoreId', question.score_id,
                               'answers', array_to_json(array(
                                SELECT json_build_object(
                                               'id', answer_id,
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
	var rows uint8
	slice := reflect.New(reflect.SliceOf(reflect.TypeOf(rows)))

	err := tx.Get(slice.Interface(), q, ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return slice.Elem().Bytes(), nil
}

func (s *TestService) Create(tx *sqlx.Tx, test models.TestRequest) (*int, error) {
	var id int

	err := tx.QueryRowx(`CALL insert_test($1, $2, $3, $4)`, test.Name, pq.Array(test.QuestionIDs), pq.Array(test.TypeIDs), id).Scan(&id)
	if err != nil {
		tx.Rollback()
		return nil, middleware.BadRequest.AddError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &id, nil
}
