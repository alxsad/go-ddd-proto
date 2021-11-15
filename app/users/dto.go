package users

import "app/domain/user"

type ListUsersQuery struct {
}

type CreateUserCmd struct {
	Name  string
	Email string
}

func (cmd CreateUserCmd) ToUserDTO() UserDTO {
	return UserDTO{
		Name:  cmd.Name,
		Email: cmd.Email,
	}
}

type UserDTO struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func (dto UserDTO) FromUserEntity(user user.UserEntity) UserDTO {
	dto.UserID = user.ID.String()
	dto.Name = user.Name
	dto.Email = user.Email.String()
	return dto
}
