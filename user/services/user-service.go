package userservice

import (
	"encoding/json"
	"fmt"
	"log"

	server "github.com/pooya/delivery/rabbit-mq"
	"github.com/pooya/entity"
	"github.com/pooya/repository"
	"github.com/streadway/amqp"
)

type Service struct {
	Repo         repository.Repo
	RabbitMqConn *server.RabbitMQ
}

func New(repo repository.Repo, conn *server.RabbitMQ) (Service, error) {
	return Service{Repo: repo, RabbitMqConn: conn}, nil
}

func (r *Service) SendExpense(user entity.User, rabbitmq_queue string) error {

	q, err := r.RabbitMqConn.Channel.QueueDeclare(
		rabbitmq_queue, // queue name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	body, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	err = r.RabbitMqConn.Channel.Publish(
		"",     // exchange
		q.Name, // routing key (queue name)
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	log.Printf("Coffee order has been sent to RabbitMQ queue: %s", user)

	return nil
}

func (s *Service) AddExpenseHandler() (entity.User, error) {
	fmt.Println("adding")

	user := entity.User{
		Name:        "pooya",
		PhoneNumber: "09202230930",
	}

	log.Println(user)

	s.Repo.AddExpense(100)

	s.SendExpense(user, "expense")

	log.Printf("adding ends")

	return user, nil

}
