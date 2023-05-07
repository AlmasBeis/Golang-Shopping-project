package service

import (
	"final-project/pkg/repository"
)

type UserService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Deposit(id int, dep float32) error {

	err := s.repo.Deposit(id, dep)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserService) GetBalance(id int) (float32, error) {

	balance, err := s.repo.GetBalance(id)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
