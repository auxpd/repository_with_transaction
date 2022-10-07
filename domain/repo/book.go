package repo

import (
	"context"

	"repoWithTrans/domain/model"
)

type BookRepo interface {
	// WithCtx set context
	WithCtx(ctx context.Context) BookRepo
	// GetByID returns a book by id
	GetByID(id int64) (*model.Book, error)
	// GetByName returns a book by name
	GetByName(name string) (*model.Book, error)
	// Save saves a book
	Save(user *model.Book) error
	// Remove remove a book
	Remove(id int64) error
}
