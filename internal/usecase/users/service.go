package users

type Service struct {
	repository Repository
}

func LoadUserService(repository Repository) *Service {
	return &Service{repository: repository}
}
