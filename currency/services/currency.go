package services

import (
	"fmt"

	"github.com/expse/repository"
)

type Service struct {
	repo *repository.Repo
}

func NewService(currencyRepo *repository.Repo) *Service {
	srv := &Service{
		repo: currencyRepo,
	}

	return srv
}

func (s *Service) CreateCurrencyRequest() error {

	err := s.repo.CreateCurrencyRequest()

	if err != nil {
		fmt.Println("error on create expense service layer happened")
	}

	return nil
}
