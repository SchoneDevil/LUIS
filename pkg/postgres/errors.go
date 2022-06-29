package postgres

import "fmt"

var (
	ErrUnique   = fmt.Errorf("Данный Email уже используется")
	ErrNotFound = fmt.Errorf("Данный пользователь не существует")
)

func CheckError(errCode string) error {
	switch errCode {
	case "23505":
		return ErrUnique
	}
	return fmt.Errorf("")
}
