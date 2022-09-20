package repo

import (
	"context"

	"repoWithTrans/domain/model"
)

type UserRepo interface {
	// WithCtx set context
	WithCtx(ctx context.Context) UserRepo
	// GetByID returns a user by id
	GetByID(id int64) (*model.User, error)
	// GetByUsername returns a user by username
	GetByUsername(username string) (*model.User, error)
	// Save saves a user
	Save(user *model.User) error
	// Remove remove a user
	Remove(id int64) error
}
