package handler

import (
	"log"
	"net/http"

	"github.com/RaniAgus/go-starter/internal/sql"
	"github.com/RaniAgus/go-starter/internal/util"
	"github.com/RaniAgus/go-starter/templates"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Queries  sql.Querier
	Validate *validator.Validate
}

// Route handlers

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h Handler) GetHome(w http.ResponseWriter, r *http.Request) error {
	v, err := h.Queries.GetLatestVersion(r.Context())
	if err != nil {
		return err
	}

	return templates.ShowHome(v.Version).Render(r.Context(), w)
}

func (h Handler) NotFound(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotFound)
	return templates.ShowErrorPage("We couldn't find the page you were looking for").Render(r.Context(), w)
}

// Error handlers

type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

func (h Handler) HandlePageError(w http.ResponseWriter, r *http.Request, err error) {
	msg := "An error has occurred. Please try again later."
	if apiError, ok := err.(util.APIError); ok {
		msg = apiError.Message
	} else {
		log.Println(err)
	}

	templates.ShowErrorPage(msg).Render(r.Context(), w)
}
