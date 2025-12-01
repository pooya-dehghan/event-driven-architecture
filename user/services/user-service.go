package userservice

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/pooya/entity"
	"github.com/pooya/params"
	"github.com/pooya/repository"
)

type Service struct {
	Repo    repository.Repo
	NatConn *nats.Conn
}

func New(repo repository.Repo, natConn *nats.Conn) (Service, error) {
	return Service{Repo: repo, NatConn: natConn}, nil
}

func (s *Service) Register() (entity.User, error) {
	return entity.User{}, nil
}

func (s *Service) AddCurrencyRequestHandler(req params.CurrencyRequestParams) (entity.User, error) {
	fmt.Println("adding")

	user := entity.User{
		Name:        "pooya",
		PhoneNumber: "09202230930",
	}

	log.Println(user)

	s.Repo.AddExpense(100)

	currencyReq := params.CurrencyRequestParams{
		UserID:      req.UserID,
		Price:       req.Price,
		Description: req.Description,
	}

	data, err := json.Marshal(currencyReq)
	if err != nil {
		log.Println("Error marshaling currency request:", err)
	}

	if err := s.NatConn.Publish("add-currency-request", data); err != nil {
		log.Println("Error publishing to NATS:", err)
	}

	log.Printf("adding ends")

	return user, nil

}
