package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"tutorial.sqlc.dev/app/tutorial"
)

func run() error {
	ctx := context.Background()

	secret := os.Getenv("POSTGRES_SECRET")
	if secret == "" {
		return fmt.Errorf("failed to read POSTGRES_SECRET from environment.")
	}
	connStr := fmt.Sprintf("postgres://postgres:%s@127.0.0.1:5432/tutorial?sslmode=disable", secret)
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
