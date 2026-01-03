package server

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var RabbitMQClient *RabbitMQ

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQConnection() {

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONNECTION_URL"))
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
}

func (r *RabbitMQ) ConsumeRabbitMQQueue(queue_name string) (<-chan amqp.Delivery, error) {

	q, err := r.Channel.QueueDeclare(
		queue_name, // name of the queue
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)

	if err != nil {
		log.Printf("Failed to declare a RabbitMQ queue: %s", err)
		return nil, err
	}

	msgs, err := r.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Printf("Failed to register a RabbitMQ consumer: %s", err)
		return nil, err
	}

	return msgs, nil
}

func (r *RabbitMQ) CloseConnection() {
	r.Channel.Close()
	r.Conn.Close()
}
