package create

import (
	"Api/internal/domains/models"
	"Api/internal/middleware"
	"Api/internal/repository/answer"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type createAnswerQuery struct {
	Answer models.Answer
}

func (q *createAnswerQuery) Validate() error {
	if q.Answer.Text == "" {
		return middleware.BadRequest
	}

	return nil
}

func fromRequest(r *http.Request) (*createAnswerQuery, error) {
	q := &createAnswerQuery{}
	var answer models.Answer

	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		return nil, err
	}
	fmt.Println(answer)

	q.Answer = answer
	return q, nil
}

func CreateAnswer(db *sqlx.DB, rep answer.Repository) http.HandlerFunc {
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
		id, err := rep.Create(tr, query.Answer)
		if err != nil {
			return nil
		}

		err = json.NewEncoder(w).Encode(id)
		if err != nil {
			return err
		}

		return nil
	})
}
