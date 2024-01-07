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
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Get("/", Route(h.GetHome, h.HandlePageError))

	return r
}

func Route(handlerFn web.HandlerFunc, errorFn web.ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFn(w, r); err != nil {
			errorFn(w, r, err)
		}
	}
}