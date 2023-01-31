package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dwiw96/learning-go/crud-api-postgres/pg/pgx"
	"github.com/dwiw96/learning-go/crud-api-postgres/pkg/models"
	//"github.com/gorilla/mux"
)

func AddTable(w http.ResponseWriter, r *http.Request) {
	pgx.RunMigrate(Ctx, DbRepo)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("table created")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var temp models.Book
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		log.Println("decoding request body to struct failed")
	}

	pgx.RunCreate(Ctx, DbRepo, temp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(temp)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	data := pgx.RunAll(Ctx, DbRepo)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	title := r.URL.Query()["title"][0]
	data := pgx.RunGetByTitle(Ctx, DbRepo, title)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var temp models.Book
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		log.Println("decoding request body failed")
	}
	id := r.URL.Query()["id"][0]
	idNew, err := strconv.Atoi(id)
	data := pgx.RunUpdate(Ctx, DbRepo, idNew, temp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query()["title"][0]
	pgx.RunDelete(Ctx, DbRepo, title)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Book Deleted")
}
