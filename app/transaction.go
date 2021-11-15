package app

import "context"

type Transaction interface {
	Finish(context.Context) error
	Rollback(context.Context) error
}

type TransactionManager interface {
	Start(context.Context) (Transaction, error)
}
