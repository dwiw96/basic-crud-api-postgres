package pgx

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
		log.Fatal(err)
	}

	fmt.Println("MIGRATE DONE!")
}

func RunCreate(ctx context.Context, pgRepo pg.Repository) {
	fmt.Println("INSERT DATA RUN")
	err := pgRepo.Create(ctx, models.Book1)
	if errors.Is(err, pg.ErrDuplicate) {
		fmt.Printf("record: %+v already exist\n", models.Book1)
	} else if err != nil {
		log.Print(err)
	}
	fmt.Println("DATA INSERTED!")
}

func RunAll(ctx context.Context, pgRepo pg.Repository) []models.Book {
	fmt.Println("GET ALL RUN")
	all, err := pgRepo.All(ctx)
	if err != nil {
		log.Print(err)
		return nil
	}

	return all
}

func RunGetByTitle(ctx context.Context, pgRepo pg.Repository) *models.Book {
	fmt.Println("GET BY TITLE RUN")
	data, err := pgRepo.GetByTitle(ctx, "Math")
	if err != nil {
		log.Print(err)
		//return nil
	}
	return data
}

func RunUpdate(ctx context.Context, pgRepo pg.Repository) *models.Book {
	fmt.Println("UPDATE RUN")
	//var updateRecord = models.Book{Release: 2022}
	data, err := pgRepo.Update(ctx, 1, models.Book2)
	if err != nil {
		log.Print(err)
		return nil
	}
	return data
}

func RunDelete(ctx context.Context, pgRepo pg.Repository) {
	fmt.Println("DELETE RUN")
	err := pgRepo.Delete(ctx, "Math")
	if err != nil {
		log.Print(err)
	}
}
