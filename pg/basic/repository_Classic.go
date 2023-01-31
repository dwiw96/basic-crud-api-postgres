package basic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dwiw96/learning-go/crud-api-postgres/pg"
	"github.com/dwiw96/learning-go/crud-api-postgres/pkg/models"

	"github.com/jackc/pgconn"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Migrate(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS books(
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL UNIQUE,
		author TEXT NOT NULL,
		release INT NOT NULL);
	`

	_, err := r.db.ExecContext(ctx, query)
	return err
}

func (r *PostgresRepository) Create(ctx context.Context, book models.Book) (*models.Book, error) {
	var res models.Book
	err := r.db.QueryRowContext(ctx, "INSERT INTO books(title, author, release) values($1, $2, $3) RETURNING id, title, author, release", book.Title, book.Author, book.Release).Scan(&res.ID, &res.Title, &res.Author, &res.Release)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, pg.ErrDuplicate
			}
		}
		return nil, err
	}
	return &res, nil
}

func (r *PostgresRepository) All(ctx context.Context) ([]models.Book, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Release); err != nil {
			return nil, err
		}
		all = append(all, book)
	}
	return all, nil
}

func (r *PostgresRepository) GetByTitle(ctx context.Context, name string) (*models.Book, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM books WHERE title=$1", name)

	var book models.Book
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Release); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pg.ErrNotExist
		}
		return nil, err
	}

	return &book, nil
}

func (r *PostgresRepository) Update(ctx context.Context, id int, updated models.Book) (*models.Book, error) {
	res, err := r.db.ExecContext(ctx, "UPDATE books SET title = $1, author = $2, Release = $3 WHERE id = $4", updated.Title, updated.Author, updated.Release, id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, pg.ErrDuplicate
			}
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, pg.ErrUpdateFailed
	}

	return &updated, nil
}

func (r *PostgresRepository) Delete(ctx context.Context, name string) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM books WHERE title=$1", name)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return pg.ErrDeleteFailed
	}
	return nil
}
