package handler

import (
	"net/http"

	"github.com/RaniAgus/go-starter/internal/sql"
	"github.com/RaniAgus/go-starter/templates"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Queries  sql.Querier
	Validate *validator.Validate
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h Handler) GetHome(w http.ResponseWriter, r *http.Request) error {
	v, err := h.Queries.GetLatestVersion(r.Context())
	if err != nil {
		return err
	}

	return templates.ShowHome(v.Version).Render(r.Context(), w)
}

func (h Handler) NotFound(w http.ResponseWriter, r *http.Request) error {
	msg := "We couldn't find the page you were looking for"
	return templates.ShowErrorPage(msg).Render(r.Context(), w)
}
