package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dwiw96/learning-go/crud-api-postgres/pg"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	dbRepo *pg.PostgresRepository
	ctx    context.Context
	cancel context.CancelFunc
)

func Book(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		pg.RunCreate(ctx, dbRepo)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Data inserted")
	case "GET":
		all := pg.RunAll(ctx, dbRepo)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(all)
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
	dbRepo = pg.NewPostgresRepository(db)

	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	pg.RunMigrate(ctx, dbRepo)

	http.HandleFunc("/", Book)

	log.Println("Listening on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
