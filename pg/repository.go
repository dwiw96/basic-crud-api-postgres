package pg

import (
	"context"
	"errors"

	"github.com/dwiw96/learning-go/crud-api-postgres/pkg/models"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("record does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type Repository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, books models.Book) error
	All(ctx context.Context) ([]models.Book, error)
	GetByTitle(ctx context.Context, name string) (*models.Book, error)
	Update(ctx context.Context, id int, updated models.Book) (*models.Book, error)
	Delete(ctx context.Context, name string) error
}
