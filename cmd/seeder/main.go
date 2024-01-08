package main

import (
	"context"
	"log"

	"github.com/RaniAgus/go-starter/data"
	"github.com/RaniAgus/go-starter/data/sqlc"
)

func main() {
	db := data.Connect()
	defer db.Close(context.Background())

	q := sqlc.New(db)
	films, err := q.ListFilms(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, film := range films {
		q.DeleteFilm(context.Background(), film.ID)
	}

	q.CreateFilms(context.Background(), []sqlc.CreateFilmsParams{
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
		{Title: "The Dark Knight", Director: "Christopher Nolan"},
		{Title: "Back to the Future", Director: "Robert Zemeckis"},
	})
}
