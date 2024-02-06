package get

import (
	"Api/internal/middleware"
	"Api/internal/repository/answer"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type getAnswerQuery struct {
	Id int64
}

func (q *getAnswerQuery) Validate() error {
	if q.Id <= 0 {
		return middleware.BadRequest
	}

	return nil
}

func fromRequest(r *http.Request) (*getAnswerQuery, error) {
	request := &getAnswerQuery{}

	param := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return nil, middleware.BadRequest
	}

	request.Id = id

	return request, nil
}

func GetAnswer(db *sqlx.DB, rep answer.Repository) http.HandlerFunc {
	return middleware.ErrorMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		query, err := fromRequest(r)
		if err != nil {
			return err
		}

		err = query.Validate()
		if err != nil {
			return err
		}

		tr := db.MustBegin()
		answer, err := rep.GetById(tr, query.Id)
		if err != nil {
			return err
		}

		err = json.NewEncoder(w).Encode(answer)
		if err != nil {
			return err
		}

		return nil
	})
}
