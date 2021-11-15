package repository

import (
	"app/domain/user"
	"app/infra/transaction"
	"context"
	"fmt"
)

type DbUsersRepository struct {
}

func NewDbUsersRepository() DbUsersRepository {
	return DbUsersRepository{}
}

func (repo DbUsersRepository) ListUsers(ctx context.Context) (out user.Users, err error) {
	var tx transaction.DbTransaction
	if tx, err = repo.getTx(ctx); err != nil {
		return out, err
	}
	var sql = "SELECT * FROM users"
	err = tx.Tx.SelectContext(ctx, &out, sql)
	if err != nil {
		return out, fmt.Errorf("sql err: %w", err)
	}
	return
}

func (repo DbUsersRepository) CreateUser(ctx context.Context, entity user.UserEntity) (err error) {
	var tx transaction.DbTransaction
	if tx, err = repo.getTx(ctx); err != nil {
		return err
	}
	var sql = "INSERT INTO users (id, name, email) VALUES (:id, :name, :email) RETURNING *"
	_, err = tx.Tx.NamedExecContext(ctx, sql, entity)
	if err != nil {
		return fmt.Errorf("sql err: %w", err)
	}
	return
}

func (repo DbUsersRepository) getTx(ctx context.Context) (tx transaction.DbTransaction, err error) {
	var ok bool
	if tx, ok = ctx.Value("tx").(transaction.DbTransaction); !ok {
		return tx, fmt.Errorf("tx can not be found")
	}
	return
}
