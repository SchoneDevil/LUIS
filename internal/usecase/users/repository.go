package users

import "user_service/internal/entity"

type Repository interface {
	GetUser(id string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUsers() []entity.User
	AddUser(user CreateUserDTO) error
	UpdateUser(id string, user UpdateUserDTO) error
	DeleteUser(id string) error
}
