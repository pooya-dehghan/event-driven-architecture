package services

import (
	"fmt"

	"github.com/expse/params"
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

func (s *Service) CreateCurrencyRequest(params params.CurrencyRequestParams) error {

	err := s.repo.CreateCurrencyRequest(params)

	if err != nil {
		fmt.Println("error on create expense service layer happened")
	}

	return nil
}
