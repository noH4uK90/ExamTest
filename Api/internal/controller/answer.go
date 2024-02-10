package controller

import (
	"net/http"
	"strconv"

	"Api/internal/domains/models"
	"Api/internal/middleware"
	"Api/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type AnswerController struct{}

func NewAnswerController() *AnswerController { return &AnswerController{} }

func (c *AnswerController) Init(r chi.Router, rep repository.AnswerRepository, db *sqlx.DB) {
	r.Route("/answer", func(r chi.Router) {
		r.Get("/", getAnswers(db, rep))
		r.Get("/{id}", getAnswer(db, rep))
		r.Post("/", createAnswer(db, rep))
		r.Put("/{id}", updateAnswer(db, rep))
		r.Delete("/{id}", deleteAnswer(db, rep))
	})
}

func getAnswers(db *sqlx.DB, rep repository.AnswerRepository) http.HandlerFunc {
	return middleware.ErrorMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		tx := db.MustBegin()
		answers, err := rep.Get(tx)
		if err != nil {
			return err
		}

		render.JSON(w, r, answers)
		return nil
	})
}

func getAnswer(db *sqlx.DB, rep repository.AnswerRepository) http.HandlerFunc {

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
		answer, err := rep.GetById(tx, req.ID)
		if err != nil {
			return err
		}

		render.JSON(w, r, answer)
		return nil
	})
}

func createAnswer(db *sqlx.DB, rep repository.AnswerRepository) http.HandlerFunc {

	type Request struct {
		answer models.Answer
	}

	return middleware.ErrorMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		var req Request
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return middleware.BadRequest
		}

		if err := validator.New().Struct(req.answer); err != nil {
			return middleware.BadRequest
		}

		tx := db.MustBegin()
		id, err := rep.Create(tx, req.answer)
		if err != nil {
			return err
		}

		render.JSON(w, r, id)
		return nil
	})
}

func updateAnswer(db *sqlx.DB, rep repository.AnswerRepository) http.HandlerFunc {

	type Request struct {
		ID     int64
		answer models.Answer
	}

	return middleware.ErrorMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		var req Request

		param := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return middleware.BadRequest
		}
		req.ID = id

		if err := render.DecodeJSON(r.Body, &req.answer); err != nil {
			return middleware.BadRequest
		}

		if err := validator.New().Struct(req.answer); err != nil {
			return middleware.BadRequest
		}

		tx := db.MustBegin()
		err = rep.Update(tx, req.ID, req.answer)
		if err != nil {
			return err
		}

		render.JSON(w, r, http.StatusNoContent)
		return nil
	})
}

func deleteAnswer(db *sqlx.DB, rep repository.AnswerRepository) http.HandlerFunc {

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
		err = rep.Delete(tx, req.ID)
		if err != nil {
			return err
		}

		render.JSON(w, r, http.StatusNoContent)
		return nil
	})
}
