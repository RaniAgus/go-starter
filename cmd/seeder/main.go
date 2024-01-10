package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RaniAgus/go-starter/internal"
	"github.com/RaniAgus/go-starter/internal/sql"
)

func main() {
	db := internal.NewDatabase()
	defer db.Close(context.Background())

	queries := sql.New(db)

	// Insert your seed data here
	versions, err := queries.ListVersions(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	queries.CreateVersion(context.Background(), fmt.Sprintf("v0.%d.%d", len(versions)+1, 0))
}
