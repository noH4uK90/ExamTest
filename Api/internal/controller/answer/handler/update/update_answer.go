package update

import (
	"Api/internal/domains/models"
	"Api/internal/middleware"
	"Api/internal/repository/answer"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type updateAnswerQuery struct {
	Id     int64
	answer models.Answer
}

func (q *updateAnswerQuery) Validate() error {
	if q.Id <= 0 || q.answer.Text == "" {
		return middleware.BadRequest
	}

	return nil
}

func fromRequest(r *http.Request) (*updateAnswerQuery, error) {
	q := &updateAnswerQuery{}
	var answer models.Answer

	param := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		return nil, err
	}

	q.Id = id
	q.answer = answer
	return q, nil
}

func UpdateAnswer(db *sqlx.DB, rep answer.Repository) http.HandlerFunc {
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
		err = rep.Update(tr, q.Id, q.answer)
		if err != nil {
			return err
		}

		return nil
	})
}
