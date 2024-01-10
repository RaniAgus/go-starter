package handler

import (
	"log"
	"net/http"

	"github.com/RaniAgus/go-starter/internal/util"
	"github.com/RaniAgus/go-starter/templates"
)

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
