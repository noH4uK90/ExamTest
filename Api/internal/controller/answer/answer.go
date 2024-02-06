package answer

import (
	"Api/internal/controller/answer/handler/get"
	"Api/internal/controller/answer/handler/get_list"
	"Api/internal/repository/answer"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type AnswerController struct{}

func NewAnswerController() *AnswerController { return &AnswerController{} }

func (c *AnswerController) Init(r *chi.Mux, rep answer.Repository, db *sqlx.DB) {
	r.Route("/answer", func(r chi.Router) {
		r.Get("/", get_list.GetAnswers(db, rep))
		r.Get("/{id}", get.GetAnswer(db, rep))
	})
}
