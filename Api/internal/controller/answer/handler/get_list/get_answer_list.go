package get_list

import (
	"Api/internal/middleware"
	"Api/internal/repository/answer"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func GetAnswers(db *sqlx.DB, rep answer.Repository) http.HandlerFunc {
	return middleware.ErrorMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		tr := db.MustBegin()
		answers, err := rep.Get(tr)
		if err != nil {
			return err
		}

		err = json.NewEncoder(w).Encode(answers)
		if err != nil {
			return err
		}

		return nil
	})
}
