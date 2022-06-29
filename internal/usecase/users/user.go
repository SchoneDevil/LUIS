package users

import (
	"time"

	"user_service/internal/entity"
	psw "user_service/pkg/password"

	"github.com/dgrijalva/jwt-go"
)

func (srv Service) Login(email, password string) (token string, err error) {
	details, err := srv.repository.GetUserByEmail(email)
	if err != nil {
		return token, entity.ErrUserDoesNotExist
	}
	if !(psw.CheckHash(details.Password, password)) {
		return token, entity.ErrInvalidPassword
	}
	//TODO move to controller and switch return to user_id
	token, err = CreateToken(details.ID)

	if err != nil {
		return token, err
	}
	return token, nil
}
func (srv Service) CreateUser(dto CreateUserDTO) error {
	return srv.repository.AddUser(dto)
}
func (srv Service) UpdateUser(id string, dto UpdateUserDTO) error {
	return srv.repository.UpdateUser(id, dto)
}
func (srv Service) DeleteUser(id string) error {
	return srv.repository.DeleteUser(id)
}
func (srv Service) GetUser(id string) (entity.User, error) {
	return srv.repository.GetUser(id)
}
func (srv Service) GetUserByEmail(email string) (entity.User, error) {
	return srv.repository.GetUserByEmail(email)
}
func (srv Service) GetUsers() []entity.User {
	return srv.repository.GetUsers()
}

func CreateToken(id string) (string, error) {
	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	now := time.Now().Local()
	token.Claims = jwt.MapClaims{
		"id":  id,
		"iat": now.Unix(),
		"exp": now.Add(time.Hour * time.Duration(1)).Unix(),
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("t0k3n"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
