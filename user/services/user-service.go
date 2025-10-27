package services

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/pooya/entity"
	"github.com/pooya/repository"
)

type UserService struct {
	Repo    repository.Repo
	NatConn *nats.Conn
}

func New(repo repository.Repo) (UserService, error) {
	return UserService{Repo: repo}, nil
}

func (s *UserService) AddExpenseHandler() (entity.User, error) {
	fmt.Println("adding")

	user := entity.User{
		Id:          "1",
		Name:        "pooya",
		PhoneNumber: "09202230930",
	}

	log.Println(user)

	if err := s.NatConn.Publish("add-expense", []byte("New Expense Added")); err != nil {
		log.Println("Error publishing to NATS:", err)
	}

	log.Printf("adding ends")

	return user, nil

}
