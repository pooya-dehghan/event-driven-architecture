package main

import (
	"encoding/json"
	"fmt"
	"log"

	server "github.com/expse/delivery"
)

func init() {

	server.NewRabbitMQConnection()

}

func main() {
	defer server.RabbitMQClient.CloseConnection()

	msgs, err := server.RabbitMQClient.ConsumeRabbitMQQueue("expense")

	if err != nil {
		log.Fatalf("Failed to consume RabbitMQ queue: %s", err)
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var coffeeOrder CoffeeOrder
			err := json.Unmarshal(d.Body, &coffeeOrder)
			if err != nil {
				log.Printf("Error reading coffee order (please check the JSON format): %s", err)
				continue
			}

			fmt.Printf("Received a coffee order: Coffee Type = %s, Price = %f\n", coffeeOrder.CoffeeType, coffeeOrder.Price)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
