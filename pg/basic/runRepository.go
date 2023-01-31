package basic

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/dwiw96/learning-go/crud-api-postgres/pg"
	"github.com/dwiw96/learning-go/crud-api-postgres/pkg/models"
)

func RunMigrate(ctx context.Context, pgRepo pg.Repository) {
	fmt.Println("MIGRATE RUN")

	if err := pgRepo.Migrate(ctx); err != nil {
		log.Println(err)
	}

	fmt.Println("MIGRATE DONE!")
}

func RunCreate(ctx context.Context, pgRepo pg.Repository, book models.Book) *models.Book {
	fmt.Println("INSERT DATA RUN")
	res, err := pgRepo.Create(ctx, book)
	if errors.Is(err, pg.ErrDuplicate) {
		fmt.Printf("record: %+v already exist\n", book)
	} else if err != nil {
		log.Println(err)
	}
	fmt.Println("DATA INSERTED!")
	return res
}

func RunAll(ctx context.Context, pgRepo pg.Repository) []models.Book {
	fmt.Println("GET ALL RUN")
	all, err := pgRepo.All(ctx)
	if err != nil {
		log.Println(err)
	}

	return all
}

func RunGetByTitle(ctx context.Context, pgRepo pg.Repository, title string) *models.Book {
	fmt.Println("GET BY TITLE RUN")
	data, err := pgRepo.GetByTitle(ctx, title)
	if err != nil {
		log.Println(err)
	}
	return data
}

func RunUpdate(ctx context.Context, pgRepo pg.Repository, id int, book models.Book) *models.Book {
	fmt.Println("UPDATE RUN")
	data, err := pgRepo.Update(ctx, id, book)
	if err != nil {
		log.Println(err)
	}
	return data
}

func RunDelete(ctx context.Context, pgRepo pg.Repository, title string) {
	fmt.Println("DELETE RUN")
	err := pgRepo.Delete(ctx, title)
	if err != nil {
		log.Println(err)
	}
}
