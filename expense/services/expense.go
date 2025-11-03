package services

import (
	"fmt"

	"github.com/expse/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	repo *repository.Repository
}

func NewService(expenseRepo *repository.Repository) *Service {
	srv := &Service{
		repo: expenseRepo,
	}

	return srv
}

func (s *Service) addExpense() (*mongo.InsertOneResult, error) {
	exp, err := s.repo.CreateExpense()

	if err != nil {
		fmt.Println("error on create expense service layer happened")
	}

	return exp, nil
}
