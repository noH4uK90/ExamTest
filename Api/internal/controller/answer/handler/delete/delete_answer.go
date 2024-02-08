package delete

import (
	"Api/internal/middleware"
	"Api/internal/repository/answer"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type deleteAnswerQuery struct {
	Id int64
}

func (q *deleteAnswerQuery) Validate() error {
	if q.Id <= 0 {
		return middleware.BadRequest
	}

	return nil
}

func fromRequest(r *http.Request) (*deleteAnswerQuery, error) {
	q := &deleteAnswerQuery{}

	param := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return nil, middleware.BadRequest
	}

	q.Id = id
	return q, nil
}

func DeleteAnswer(db *sqlx.DB, rep answer.Repository) http.HandlerFunc {
	return middleware.ErrorMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		q, err := fromRequest(r)
		if err != nil {
			return err
		}

		err = q.Validate()
		if err != nil {
			return err
		}

		tr := db.MustBegin()
		err = rep.Delete(tr, q.Id)
		if err != nil {
			return err
		}

		return nil
	})
}
