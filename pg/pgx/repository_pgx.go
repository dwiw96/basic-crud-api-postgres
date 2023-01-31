package pgx

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dwiw96/learning-go/crud-api-postgres/pg"
	"github.com/dwiw96/learning-go/crud-api-postgres/pkg/models"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresRepositoryPGX struct {
	db *pgxpool.Pool
}

func NewPostgresRepositoryPGX(db *pgxpool.Pool) *PostgresRepositoryPGX {
	return &PostgresRepositoryPGX{
		db: db,
	}
}

func (r *PostgresRepositoryPGX) Migrate(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS books(
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL UNIQUE,
		author TEXT NOT NULL,
		release INT NOT NULL);
		`
	_, err := r.db.Exec(ctx, query)
	return err
}

func (r *PostgresRepositoryPGX) Create(ctx context.Context, book models.Book) (*models.Book, error) {
	var res models.Book
	err := r.db.QueryRow(ctx, "INSERT INTO books(title, author, release) VALUES($1, $2, $3) RETURNING id, title, author, release", book.Title, book.Author, book.Release).Scan(&res.ID, &res.Title, &res.Author, &res.Release)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, pg.ErrDuplicate
			}
		}
		return nil, err
	}
	return &res, err
}

func (r *PostgresRepositoryPGX) All(ctx context.Context) ([]models.Book, error) {
	res, err := r.db.Query(ctx, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var all []models.Book
	for res.Next() {
		var book models.Book
		if err := res.Scan(&book.ID, &book.Title, &book.Author, &book.Release); err != nil {
			return nil, err
		}
		all = append(all, book)
	}
	return all, nil
}

func (r *PostgresRepositoryPGX) GetByTitle(ctx context.Context, title string) (*models.Book, error) {
	row := r.db.QueryRow(ctx, "SELECT * FROM books WHERE title=$1", title)

	var book models.Book
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Release); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pg.ErrNotExist
		}
		return nil, err
	}

	return &book, nil
}

func (r *PostgresRepositoryPGX) Update(ctx context.Context, id int, book models.Book) (*models.Book, error) {
	res, err := r.db.Exec(ctx, "UPDATE books SET title=$1, author=$2, release=$3 WHERE id=$4", book.Title, book.Author, book.Release, id)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, pg.ErrDuplicate
			}
		}
		return nil, err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, pg.ErrUpdateFailed
	}

	book.ID = id
	return &book, nil
}

func (r *PostgresRepositoryPGX) Delete(ctx context.Context, name string) error {
	res, err := r.db.Exec(ctx, "DELETE FROM books WHERE title=$1", name)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return pg.ErrDeleteFailed
	}
	return nil
}
