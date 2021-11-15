package app

import (
	"app/app/users"
	"app/domain/user"
	"context"
	"fmt"
)

func (app *Application) ListUsers(ctx context.Context, qry users.ListUsersQuery) (list []users.UserDTO, err error) {
	var tx Transaction
	if tx, err = app.tm.Start(ctx); err != nil {
		return
	}
	ctx = context.WithValue(ctx, "tx", tx)

	var userEntities user.Users
	userEntities, err = app.users.ListUsers(ctx)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return list, rollbackErr
		}
		return
	}

	if err = tx.Finish(ctx); err != nil {
		fmt.Errorf("can not commit tx")
		return
	}

	for _, entity := range userEntities {
		list = append(list, users.UserDTO{}.FromUserEntity(entity))
	}

	return
}

func (app *Application) CreateUser(ctx context.Context, cmd users.CreateUserCmd) (dto users.UserDTO, err error) {
	var tx Transaction
	if tx, err = app.tm.Start(ctx); err != nil {
		return
	}
	ctx = context.WithValue(ctx, "tx", tx)

	var entity user.UserEntity
	if entity, err = app.users.CreateUser(ctx, cmd.ToUserDTO()); err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return dto, rollbackErr
		}
		return
	}

	if err = tx.Finish(ctx); err != nil {
		fmt.Errorf("can not commit tx")
		return
	}

	return dto.FromUserEntity(entity), nil
}
