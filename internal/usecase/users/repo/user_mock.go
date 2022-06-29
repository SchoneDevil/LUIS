package repo

import (
	"fmt"
	"time"

	"user_service/internal/entity"
	"user_service/internal/usecase/users"
	"user_service/pkg/password"
)

type MockDB struct {
}

func NewMockDB() *MockDB {
	return &MockDB{}
}

func (m MockDB) GetUserByEmail(email string) (entity.User, error) {
	users := []entity.User{
		{
			ID:       "be811bfb-ea85-405e-975f-df40dae20bad",
			Username: "test user",
			Email:    "test@example.com",
			Password: "$2a$14$sZUHOMhb/Ja.wMRhyFi1/eMxorLjoJYcQgiq2i0opjP1/nPpTLkoi", //content
		},
	}
	for _, a := range users {
		fmt.Printf("%v, %s\n", a.Email, email)
		if a.Email != email {
			return entity.User{}, fmt.Errorf("user already exists")
		}
		return a, nil
	}
	return entity.User{}, fmt.Errorf("exit circle")
}

func (m MockDB) GetUser(id string) (entity.User, error) {
	users := []entity.User{
		{
			ID:       "be811bfb-ea85-405e-975f-df40dae20bad",
			Username: "test user",
			Email:    "test@example.com",
			Password: "$2a$14$sZUHOMhb/Ja.wMRhyFi1/eMxorLjoJYcQgiq2i0opjP1/nPpTLkoi", //content
		},
	}
	for _, a := range users {
		if a.ID != id {
			return entity.User{}, fmt.Errorf("user already exists")
		}
		return a, nil
	}
	return entity.User{}, fmt.Errorf("exit circle")
}

func (m MockDB) GetUsers() []entity.User {
	users := []entity.User{
		{
			ID:       "be811bfb-ea85-405e-975f-df40dae20bad",
			Username: "test user",
			Email:    "test@example.com",
			Password: "$2a$14$sZUHOMhb/Ja.wMRhyFi1/eMxorLjoJYcQgiq2i0opjP1/nPpTLkoi", //content
		},
		{
			ID:       "37ef2e0c-2320-4bbd-91c0-647957bda0b1",
			Username: "test user 2",
			Email:    "test2@example.com",
			Password: "$2a$14$sZUHOMhb/Ja.wMRhyFi1/eMxorLjoJYcQgiq2i0opjP1/nPpTLkoi", //content
		},
	}
	return users
}

func (m MockDB) AddUser(user users.CreateUserDTO) error {
	users := []entity.User{
		{
			ID:       "",
			Username: "test user",
			Email:    "test@example.com",
			Password: "password",
		},
	}
	startCount := len(users)
	hash, err := password.Hash(user.Password)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	users = append(users, entity.User{
		ID:        "dd2a571c-c49e-4481-8984-760ffce38410",
		Username:  user.Username,
		Email:     user.Email,
		Password:  hash,
		CreatedAt: int(time.Now().UnixNano()),
		UpdatedAt: int(time.Now().UnixNano()),
	})

	if startCount == len(users) {
		return fmt.Errorf("user not add")
	}
	return nil
}

func (m MockDB) UpdateUser(id string, user users.UpdateUserDTO) error {
	users := []entity.User{
		{
			ID:       "be811bfb-ea85-405e-975f-df40dae20bad",
			Username: "test user",
			Email:    "test@example.com",
			Password: "$2a$14$sZUHOMhb/Ja.wMRhyFi1/eMxorLjoJYcQgiq2i0opjP1/nPpTLkoi", //content
		},
	}
	for _, a := range users {
		if a.ID != id {
			return fmt.Errorf("user already exists")
		}
		a.Username = user.Username
		a.Email = user.Email
		if a.Email != user.Email && a.Username != user.Username {
			return fmt.Errorf("user update faild")
		}
		return nil
	}
	return nil
}

func (m MockDB) DeleteUser(id string) error {
	users := []entity.User{
		{
			ID:       "be811bfb-ea85-405e-975f-df40dae20bad",
			Username: "test user",
			Email:    "test@example.com",
			Password: "$2a$14$sZUHOMhb/Ja.wMRhyFi1/eMxorLjoJYcQgiq2i0opjP1/nPpTLkoi", //content
		},
		{
			ID:       "37ef2e0c-2320-4bbd-91c0-647957bda0b1",
			Username: "test user 2",
			Email:    "test2@example.com",
			Password: "$2a$14$sZUHOMhb/Ja.wMRhyFi1/eMxorLjoJYcQgiq2i0opjP1/nPpTLkoi", //content
		},
	}

	for k, a := range users {
		if a.ID != id {
			return fmt.Errorf("user already exists")
		}
		newUsers := RemoveCopy(users, k)
		if len(users) == len(newUsers) {
			return fmt.Errorf("user not deleted")
		}
		return nil
	}
	return nil
}

func RemoveCopy(slice []entity.User, i int) []entity.User {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
