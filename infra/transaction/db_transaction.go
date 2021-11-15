package transaction

import (
	"app/app"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DbTransaction struct {
	Tx *sqlx.Tx
}

func (tx DbTransaction) Finish(ctx context.Context) (err error) {
	if err = tx.Tx.Commit(); err != nil {
		return fmt.Errorf("can not commit tx: %w", err)
	}
	return
}

func (tx DbTransaction) Rollback(ctx context.Context) (err error) {
	if err = tx.Tx.Rollback(); err != nil {
		return fmt.Errorf("can not rollback tx: %w", err)
	}
	return
}


type DbTransactionManager struct {
	db *sqlx.DB
}

func NewDbTransactionManager(db *sqlx.DB) DbTransactionManager {
	return DbTransactionManager{
		db: db,
	}
}

func (tm DbTransactionManager) Start(ctx context.Context) (app.Transaction, error) {
	var tx, err = tm.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, fmt.Errorf("can not start tx: %w", err)
	}
	return DbTransaction{Tx: tx}, nil
}
