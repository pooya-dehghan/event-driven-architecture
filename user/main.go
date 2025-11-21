package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/pooya/config"
	httpserver "github.com/pooya/delivery/http-server"
	"github.com/pooya/repository"
	userservice "github.com/pooya/services"
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
	cfg := config.Load("config.yml")

	repo := repository.NewRepo(&cfg.MysqlDatabase)

	nc, err := makeNatsConnection()

	if err != nil {
		panic("something happend on init")
	}

	userService, err := userservice.New(repo, nc)

	server := httpserver.NewServer(userService, cfg)

	server.Serve()

	if err != nil {
		panic("something happend on init")
	}
}
