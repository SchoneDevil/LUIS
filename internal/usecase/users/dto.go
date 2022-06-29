package users

type CreateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
