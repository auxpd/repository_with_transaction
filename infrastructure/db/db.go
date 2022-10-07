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

// GetEngine get your xorm object
func GetEngine() *gorm.DB {
	// init your db object

	return &gorm.DB{}
}

func (d GormOpsImpl) getDbFromCtx(ctx context.Context) (*gorm.DB, error) {
	if v := ctx.Value(dbFlag); v != nil {
		if db, ok := v.(*gorm.DB); ok {
			return db, nil
		}
	}

	return nil, ErrDbNotFound
}

// GetDBFromCtx try to get DB object from the context, if not found, it attempts to return defaultDB
func GetDBFromCtx(ctx context.Context, defaultDB ...*gorm.DB) *gorm.DB {
	if v := ctx.Value(dbFlag); v != nil {
		if db, ok := v.(*gorm.DB); ok {
			return db
		}
	}

	var db *gorm.DB
	if len(defaultDB) > 0 && defaultDB[0] != nil {
		db = defaultDB[0].WithContext(ctx)
	}

	return db
}

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
