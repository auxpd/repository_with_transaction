package db

import "context"

type IdbOps interface {
	// Transaction start a transaction as a block
	Transaction(ctx context.Context, fc func(ctx context.Context) error, opts ...*TxOptions) (err error)
	// Begin start a transaction
	Begin(ctx context.Context, opts ...*TxOptions) context.Context
	// Commit a transaction
	Commit(ctx context.Context) error
	// Rollback a transaction
	Rollback(ctx context.Context) error
	// Ctx get context with a DB object
	Ctx(ctx context.Context) context.Context
}
