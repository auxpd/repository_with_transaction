package services

import (
	"context"

	"repoWithTrans/domain/model"
	"repoWithTrans/domain/repo"
	"repoWithTrans/infrastructure/db"
)

// UserService profile query bus
type UserService struct {
	ops      db.IdbOps
	userRepo repo.UserRepo
}

// New create bus instance
func New(userRepo repo.UserRepo, ops db.IdbOps) *UserService {
	return &UserService{
		ops:      ops,
		userRepo: userRepo,
	}
}

func (s *UserService) Create(user *model.User) error {
	// start transaction
	err := s.ops.Transaction(context.Background(), func(ctx context.Context) (err error) {
		err = s.userRepo.WithCtx(ctx).Save(user)
		// do other repo operations

		return
	})
	if err != nil {
		return err
	}

	return nil
}
