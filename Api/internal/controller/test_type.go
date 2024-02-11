package controller

import (
	"net/http"
	"strconv"

	"Api/internal/middleware"
	"Api/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
)

type TestTypeController struct{}

func NewTestTypeController() *TestTypeController { return &TestTypeController{} }

func (c *TestTypeController) Init(r chi.Router, rep repository.TestTypeRepository, db *sqlx.DB) {
	r.Route("/testType", func(r chi.Router) {
		r.Get("/", getTestTypes(db, rep))
		r.Get("/{id}", getTestType(db, rep))
	})
}

func getTestTypes(db *sqlx.DB, rep repository.TestTypeRepository) http.HandlerFunc {
	return middleware.ErrorMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		tx := db.MustBegin()
		testTypes, err := rep.Get(tx)
		if err != nil {
			return err
		}

		render.JSON(w, r, testTypes)
		return nil
	})
}

func getTestType(db *sqlx.DB, rep repository.TestTypeRepository) http.HandlerFunc {

	type Request struct {
		ID int64
	}

	return middleware.ErrorMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		var req Request

		param := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return middleware.BadRequest
		}
		req.ID = id

		tx := db.MustBegin()
		testType, err := rep.GetById(tx, req.ID)
		if err != nil {
			return err
		}

		render.JSON(w, r, testType)
		return nil
	})
}
