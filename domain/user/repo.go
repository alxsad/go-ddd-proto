package user

import (
	"context"
)

type UsersRepository interface {
	ListUsers(ctx context.Context) (Users, error)
	CreateUser(ctx context.Context, user UserEntity) error
}
