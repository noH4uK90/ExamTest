package controller

import (
	"Api/internal/middleware"
	"Api/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type TestController struct{}

func NewTestController() *TestController { return &TestController{} }

func (c *TestController) Init(r chi.Router, rep repository.TestRepository, db *sqlx.DB) {
	r.Route("/test", func(r chi.Router) {
		r.Get("/", getTests(db, rep))
		r.Get("/{id}", getTest(db, rep))
	})
}

func getTests(db *sqlx.DB, rep repository.TestRepository) http.HandlerFunc {
	return middleware.ErrorMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		tx := db.MustBegin()
		tests, err := rep.Get(tx)
		if err != nil {
			return middleware.BadRequest.AddError(err)
		}

		render.Data(w, r, tests)
		return nil
	})
}

func getTest(db *sqlx.DB, rep repository.TestRepository) http.HandlerFunc {
	return middleware.ErrorMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		param := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return middleware.BadRequest
		}

		tx := db.MustBegin()
		test, err := rep.GetById(tx, id)

		render.Data(w, r, test)
		return nil
	})
}
