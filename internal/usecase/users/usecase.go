package users

import "user_service/internal/entity"

type UseCase interface {
	Login(email, password string) (entity.User, error)
	CreateUser(user CreateUserDTO) error
	UpdateUser(id string, dto UpdateUserDTO) error
	DeleteUser(id string) error
	GetUser(id string) entity.User
	GetUserByEmail(email string) (entity.User, error)
	GetUsers() []entity.User
}
