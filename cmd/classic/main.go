package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	//"github.com/dwiw96/learning-go/crud-api-postgres/pg"
	"github.com/dwiw96/learning-go/crud-api-postgres/pg/basic"
	"github.com/dwiw96/learning-go/crud-api-postgres/pkg/handlers"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	dbRepo *basic.PostgresRepository
	ctx    context.Context
	cancel context.CancelFunc
)

func main() {
	db, err := sql.Open("pgx", "postgres://db:secret@localhost:5432/crudAPI")
	if err != nil {
		log.Fatal(err)
	}
	if db != nil {
		log.Println("db != nil")
	}
	defer db.Close()
	//apiRepository := pg.NewPostgresRepository(db)
	dbRepo = basic.NewPostgresRepository(db)

	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	handlers.AddRepo(ctx, dbRepo)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/books", handlers.Books)

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
