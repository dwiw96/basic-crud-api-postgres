package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dwiw96/learning-go/crud-api-postgres/pg"
	"github.com/dwiw96/learning-go/crud-api-postgres/pg/basic"
	"github.com/dwiw96/learning-go/crud-api-postgres/pkg/models"
)

var (
	Ctx    context.Context
	DbRepo pg.Repository
)

func AddRepo(ctx context.Context, dbRepo pg.Repository) {
	Ctx = ctx
	DbRepo = dbRepo
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		basic.RunMigrate(Ctx, DbRepo)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("table created")
	case "GET":
		data := basic.RunAll(Ctx, DbRepo)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func Books(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var temp models.Book
	switch r.Method {
	case "POST":
		err := json.NewDecoder(r.Body).Decode(&temp)
		if err != nil {
			log.Println("Decoding json failed")
		}
		res := basic.RunCreate(Ctx, DbRepo, temp)
		json.NewEncoder(w).Encode(res)
	case "GET":
		title := r.URL.Query()["title"][0]
		res := basic.RunGetByTitle(Ctx, DbRepo, title)
		json.NewEncoder(w).Encode(res)
	case "PUT":
		id := r.URL.Query()["id"][0]
		idNew, err := strconv.Atoi(id)
		err = json.NewDecoder(r.Body).Decode(&temp)
		if err != nil {
			json.NewEncoder(w).Encode("Decoding Json failed")
		}
		res := basic.RunUpdate(Ctx, DbRepo, idNew, temp)
		json.NewEncoder(w).Encode(res)
	case "DELETE":
		title := r.URL.Query()["title"][0]
		basic.RunDelete(Ctx, DbRepo, title)
		json.NewEncoder(w).Encode("RECORD DELETED")
	}
}
