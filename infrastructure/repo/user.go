package repo

import (
	"context"

	"repoWithTrans/domain/model"
	"repoWithTrans/domain/repo"
	"repoWithTrans/infrastructure/db"
)

type UserRepoImpl struct {
	orm db.UseGorm
}

var _ repo.UserRepo = UserRepoImpl{}

func (UserRepoImpl) WithCtx(ctx context.Context) repo.UserRepo {
	return &UserRepoImpl{
		orm: db.UseGorm{Ctx: ctx},
	}
}

func (u UserRepoImpl) GetByID(id int64) (user *model.User, err error) {
	err = u.orm.DB().First(&user, id).Error
	if err != nil {
		return
	}

	return
}

func (u UserRepoImpl) GetByUsername(username string) (user *model.User, err error) {
	err = u.orm.DB().Where("username = ?", username).First(&user).Error
	if err != nil {
		return
	}

	return
}

func (u UserRepoImpl) Save(user *model.User) error {
	return u.orm.DB().Save(user).Error
}

func (u UserRepoImpl) Remove(id int64) error {
	return u.orm.DB().Where("id = ?", id).Delete(&model.User{}).Error
}
