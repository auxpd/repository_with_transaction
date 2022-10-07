package db

import (
	"context"

	"gorm.io/gorm"
)

// GormOpsImpl gormDB操作对象
type GormOpsImpl struct {
	db *gorm.DB
}

// NewGormDBOps 创建GORM DB操作对象
func NewGormDBOps(db *gorm.DB) IdbOps {
	return GormOpsImpl{db: db}
}

func (d GormOpsImpl) Transaction(ctx context.Context, fc func(ctx context.Context) error, opts ...*TxOptions) (err error) {
	session := d.db.Begin(opts...).WithContext(ctx)
	ctx = context.WithValue(ctx, dbFlag, session)

	defer func() {
		if err := recover(); err != nil {
			session.Rollback()
			panic(err)
		}
	}()
	err = fc(ctx)
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit().Error
}

func (d GormOpsImpl) Begin(ctx context.Context, opts ...*TxOptions) context.Context {
	db := d.db.WithContext(ctx).Begin(opts...)
	ctx = context.WithValue(ctx, dbFlag, db)

	return ctx
}

func (d GormOpsImpl) Commit(ctx context.Context) error {
	db, err := d.getDbFromCtx(ctx)
	if err != nil {
		return err
	}

	return db.Commit().Error
}

func (d GormOpsImpl) Rollback(ctx context.Context) error {
	db, err := d.getDbFromCtx(ctx)
	if err != nil {
		return err
	}

	return db.Rollback().Error
}

func (d GormOpsImpl) Ctx(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, dbFlag, d.db)

	return ctx
}
