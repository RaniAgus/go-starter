package web

import (
	"log"
	"net/http"
	"strings"

	"github.com/RaniAgus/go-starter/data/sqlc"
	"github.com/RaniAgus/go-starter/util"
	"github.com/RaniAgus/go-starter/web/templates"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	DB       sqlc.Querier
	Validate *validator.Validate
}

// Route handlers

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h Handler) GetHome(w http.ResponseWriter, r *http.Request) error {
	http.Redirect(w, r, "/films", http.StatusSeeOther)
	return nil
}

func (h Handler) GetFilms(w http.ResponseWriter, r *http.Request) error {
	films, err := h.DB.ListFilms(r.Context())
	if err != nil {
		return err
	}

	return templates.ShowFilms(films).Render(r.Context(), w)
}

func (h Handler) PostFilm(w http.ResponseWriter, r *http.Request) error {
	form := templates.NewFilmForm{
		Title:    strings.Trim(r.PostFormValue("title"), " "),
		Director: strings.Fields(r.PostFormValue("director")),
	}

	err := h.Validate.Struct(form)
	if err != nil {
		return templates.ShowNewFilmForm(form, util.GetValidationErrorFields(err)).Render(r.Context(), w)
	}

	film, err := h.DB.CreateFilm(r.Context(), sqlc.CreateFilmParams{
		Title:    form.Title,
		Director: strings.Join(form.Director, " "),
	})
	if err != nil {
		return err
	}

	return templates.SwapNewFilm(film).Render(r.Context(), w)
}

func (h Handler) NotFound(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotFound)
	return templates.ShowErrorPage("We couldn't find the page you were looking for").Render(r.Context(), w)
}

// Error handlers

type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

func (h Handler) HandlePageError(w http.ResponseWriter, r *http.Request, err error) {
	msg := "Something went wrong"
	if apiError, ok := err.(util.APIError); ok {
		msg = apiError.Message
	} else {
		log.Println(err)
	}

	templates.ShowErrorPage(msg).Render(r.Context(), w)
}
