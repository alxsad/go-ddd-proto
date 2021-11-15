package transaction

import (
	"app/app"
	"context"
)

type DummyTransaction struct {
}

func (tx DummyTransaction) Finish(ctx context.Context) error {
	return nil
}

func (tx DummyTransaction) Rollback(ctx context.Context) error {
	return nil
}

type DummyTransactionManager struct {
}

func NewDummyTransactionManager() DummyTransactionManager {
	return DummyTransactionManager{}
}

func (tm DummyTransactionManager) Start(ctx context.Context) (app.Transaction, error) {
	return DummyTransaction{}, nil
}
