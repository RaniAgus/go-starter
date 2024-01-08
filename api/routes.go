package api

import (
	"net/http"

	"github.com/RaniAgus/go-starter/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h web.Handler) *chi.Mux {
	r := chi.NewRouter()
	fs := http.FileServer(http.Dir("web/static"))

	r.Use(middleware.Logger)
	r.Get("/", route(h.GetHome, h.HandlePageError))
	r.Get("/", route(h.GetHome, h.HandlePageError))
	r.Get("/films", route(h.GetFilms, h.HandlePageError))
	r.Post("/films", route(h.PostFilm, h.HandlePageError))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.NotFound(route(h.NotFound, h.HandlePageError))

	return r
}

func route(handlerFn web.HandlerFunc, errorFn web.ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFn(w, r); err != nil {
			errorFn(w, r, err)
		}
	}
}
