package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dwiw96/learning-go/crud-api-postgres/pg/pgx"
	"github.com/dwiw96/learning-go/crud-api-postgres/pkg/handlers"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://db:secret@localhost:5432/crudAPI")
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	dbRepo := pgx.NewPostgresRepositoryPGX(dbpool)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	handlers.AddRepo(ctx, dbRepo)

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.AddTable).Methods(http.MethodPost)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.GetBookByTitle).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books", handlers.DeleteBook).Methods(http.MethodDelete)

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
