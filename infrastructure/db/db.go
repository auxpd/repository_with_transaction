package db

import (
	"context"
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

type dbIdentity struct{}

type (
	TxOptions      = sql.TxOptions
	IsolationLevel = sql.IsolationLevel
)

var dbFlag = dbIdentity{}

var ErrDbNotFound = errors.New("db object not found")

// UseGorm combine it so that it has DB methods
type UseGorm struct {
	Ctx context.Context
}

// DB get db object. return gorm.DB or sql.DB or whatever
func (u UseGorm) DB() *gorm.DB {
	if v := u.Ctx.Value(dbFlag); v != nil {
		if db, ok := v.(*gorm.DB); ok {
			return db
		}
	}
	return GetEngine()
}

// GetEngine get your xorm object
func GetEngine() *gorm.DB {
	// init your db object

	return &gorm.DB{}
}

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

func (d GormOpsImpl) getDbFromCtx(ctx context.Context) (*gorm.DB, error) {
	if v := ctx.Value(dbFlag); v != nil {
		if db, ok := v.(*gorm.DB); ok {
			return db, nil
		}
	}
	return nil, ErrDbNotFound
}
