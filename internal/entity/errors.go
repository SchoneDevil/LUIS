package entity

import "errors"

var ErrUserAlreadyExists = errors.New("User already exists")

var ErrUserDoesNotExist = errors.New("User does not exist")

var ErrInvalidPassword = errors.New("Invalid password")

var ErrInvalidInput = errors.New("Invalid input")
