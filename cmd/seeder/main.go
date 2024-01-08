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

	q.CreateFilm(context.Background(), sqlc.CreateFilmParams{
		Title:    "The Godfather",
		Director: "Francis Ford Coppola",
	})

	q.CreateFilm(context.Background(), sqlc.CreateFilmParams{
		Title:    "The Shawshank Redemption",
		Director: "Frank Darabont",
	})

	q.CreateFilm(context.Background(), sqlc.CreateFilmParams{
		Title:    "The Dark Knight",
		Director: "Christopher Nolan",
	})

	q.CreateFilm(context.Background(), sqlc.CreateFilmParams{
		Title:    "Back to the Future",
		Director: "Robert Zemeckis",
	})
}
