package repository

import (
	"app/domain/user"
	"context"
)

type InMemoryUsersRepository struct {
	storage *user.Users
}

func NewInMemoryUsersRepository() InMemoryUsersRepository {
	var testUser, _ = user.NewUser("John Doe", "john@doe.com")
	return InMemoryUsersRepository{
		storage: &user.Users{testUser},
	}
}

func (repo InMemoryUsersRepository) ListUsers(ctx context.Context) (user.Users, error) {
	return *(repo.storage), nil
}

func (repo InMemoryUsersRepository) CreateUser(ctx context.Context, entity user.UserEntity) error {
	*(repo.storage) = append(*(repo.storage), entity)

	return nil
}
