package userservice

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/pooya/entity"
	"github.com/pooya/repository"
)

type Service struct {
	Repo    repository.Repo
	NatConn *nats.Conn
}

func New(repo repository.Repo, natConn *nats.Conn) (Service, error) {
	return Service{Repo: repo, NatConn: natConn}, nil
}

func (s *Service) AddExpenseHandler() (entity.User, error) {
	fmt.Println("adding")

	user := entity.User{
		Name:        "pooya",
		PhoneNumber: "09202230930",
	}

	log.Println(user)

	s.Repo.AddExpense(100)

	if err := s.NatConn.Publish("add-expense", []byte("New Expense Added")); err != nil {
		log.Println("Error publishing to NATS:", err)
	}

	log.Printf("adding ends")

	return user, nil

}
