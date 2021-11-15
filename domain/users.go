package domain

import (
	"app/app/users"
	"app/domain/user"
	"context"
)

type UsersService struct {
	repo user.UsersRepository
}

func NewUsersService(repo user.UsersRepository) UsersService {
	return UsersService{
		repo: repo,
	}
}

func (s UsersService) ListUsers(ctx context.Context) (user.Users, error) {
	return s.repo.ListUsers(ctx)
}

func (s UsersService) CreateUser(ctx context.Context, dto users.UserDTO) (entity user.UserEntity, err error) {
	entity, err = user.NewUser(dto.Name, dto.Email)
	if err != nil {
		return entity, err
	}

	if err = s.repo.CreateUser(ctx, entity); err != nil {
		return entity, err
	}

	return
}
