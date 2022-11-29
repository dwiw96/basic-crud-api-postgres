package pg

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/dwiw96/learning-go/crud-api-postgres/pkg/models"
)

func RunMigrate(ctx context.Context, pgRepo Repository) {
	fmt.Println("MIGRATE RUN")

	if err := pgRepo.Migrate(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("MIGRATE DONE!")
}

func RunCreate(ctx context.Context, pgRepo Repository) {
	fmt.Println("INSERT DATA RUN")
	err := pgRepo.Create(ctx, models.Book1)
	if errors.Is(err, ErrDuplicate) {
		fmt.Printf("record: %+v already exist\n", models.Book1)
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DATA INSERTED!")
}

func RunAll(ctx context.Context, pgRepo Repository) []models.Book {
	fmt.Println("GET ALL RUN")
	all, err := pgRepo.All(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return all
}
