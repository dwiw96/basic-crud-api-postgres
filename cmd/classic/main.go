package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	//"github.com/dwiw96/learning-go/crud-api-postgres/pg"
	"github.com/dwiw96/learning-go/crud-api-postgres/pg/basic"
	//"github.com/dwiw96/learning-go/crud-api-postgres/pkg/handlers"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	dbRepo *basic.PostgresRepository
	ctx    context.Context
	cancel context.CancelFunc
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		basic.RunCreate(ctx, dbRepo)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Data inserted")
	case "GET":
		all := basic.RunAll(ctx, dbRepo)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(all)
	}
}

func book(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		data := basic.RunGetByTitle(ctx, dbRepo)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	case "PUT":
		data := basic.RunUpdate(ctx, dbRepo)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(data)
	case "DELETE":
		basic.RunDelete(ctx, dbRepo)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode("Data Deleted")
	}
}

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

	basic.RunMigrate(ctx, dbRepo)
	http.HandleFunc("/", home)
	http.HandleFunc("/book", book)

	//handlers.HandlerParam(ctx, dbRepo)
	//http.HandleFunc("/", handlers.Home)
	//http.HandleFunc("/book", handlers.Book)

	log.Println("Listening on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
