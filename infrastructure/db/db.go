package db

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type dbIdentity struct{}

type (
	TxOptions      = sql.TxOptions
	IsolationLevel = sql.IsolationLevel
)

var dbFlag = dbIdentity{}

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
