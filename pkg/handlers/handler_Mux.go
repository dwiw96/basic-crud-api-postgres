package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dwiw96/learning-go/crud-api-postgres/pg"
	"github.com/dwiw96/learning-go/crud-api-postgres/pg/pgx"
	//"github.com/dwiw96/learning-go/crud-api-postgres/models"
	//"github.com/gorilla/mux"
)

var (
	ctx    context.Context
	dbRepo pg.Repository
)

func AddRepo(time context.Context, db pg.Repository) {
	ctx = time
	dbRepo = db
}

func AddTable(w http.ResponseWriter, r *http.Request) {
	pgx.RunMigrate(ctx, dbRepo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("table created")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	pgx.RunCreate(ctx, dbRepo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Book Added")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	data := pgx.RunAll(ctx, dbRepo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	data := pgx.RunGetByTitle(ctx, dbRepo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	data := pgx.RunUpdate(ctx, dbRepo)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	pgx.RunDelete(ctx, dbRepo)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Book Deleted")
}
