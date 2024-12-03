package service

type Repository interface{}

type Service struct {
	Repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{Repository: repository}
}
