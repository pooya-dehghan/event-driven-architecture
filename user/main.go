package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	httpserver "github.com/pooya/delivery/http-server"
	"github.com/pooya/repository"
	"github.com/pooya/services"
)

func makeNatsConnection() (*nats.Conn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Drain()
	if err != nil {
		fmt.Println(err)
	}

	log.Println("nats connection was made")

	return nc, nil
}

func main() {

	repo := repository.NewRepo()
	nc, err := makeNatsConnection()
	userService, err := services.New(repo, nc)

	if err != nil {
		fmt.Errorf("something went wrong")

		return
	}

	server := httpserver.NewServer(userService)

	server.Serve()
}
