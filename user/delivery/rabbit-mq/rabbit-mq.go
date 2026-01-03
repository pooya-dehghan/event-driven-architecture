package server

import (
	"log"

	"github.com/streadway/amqp"
)

var RabbitMQClient *RabbitMQ

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQConnection() *RabbitMQ {

	conn, err := amqp.Dial("amqp://admin:admin123@localhost:5672")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a RabbitMQ channel: %s", err)
	}

	RabbitMQClient = &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}

	return RabbitMQClient
}

func (r *RabbitMQ) CloseConnection() {
	r.Channel.Close()
	r.Conn.Close()
}
