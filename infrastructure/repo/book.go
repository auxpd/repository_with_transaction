package repo

import (
	"context"

	"gorm.io/gorm"

	"repoWithTrans/domain/model"
	"repoWithTrans/domain/repo"
	"repoWithTrans/infrastructure/db"
)

type BookRepoImpl struct {
	ctx context.Context
	db  *gorm.DB
}

var _ repo.BookRepo = BookRepoImpl{}

// NewBookRepo returns a new BookRepoImpl with a default *gorm.DB
func NewBookRepo(ctx context.Context, defaultDB *gorm.DB) repo.BookRepo {
	_db := db.GetDBFromCtx(ctx, defaultDB)

	r := &BookRepoImpl{
		ctx: ctx,
		db:  _db,
	}

	return r
}

func (b BookRepoImpl) WithCtx(ctx context.Context) repo.BookRepo {
	return NewBookRepo(ctx, b.db)
}

func (b BookRepoImpl) GetByID(id int64) (book *model.Book, err error) {
	err = b.db.First(&book, id).Error
	if err != nil {
		return
	}

	return
}

func (b BookRepoImpl) GetByName(name string) (book *model.Book, err error) {
	err = b.db.Where("name = ?", name).First(&book).Error
	if err != nil {
		return
	}

	return
}

func (b BookRepoImpl) Save(book *model.Book) error {
	return b.db.Save(book).Error
}

func (b BookRepoImpl) Remove(id int64) error {
	return b.db.Where("id = ?", id).Delete(&model.Book{}).Error
}
