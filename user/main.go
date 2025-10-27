package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
	httpserver "github.com/pooya/delivery/http-server"
	"github.com/pooya/repository"
	"github.com/pooya/services"
)

type Server struct {
	natConn *nats.Conn
}

func main() {

	repo := repository.NewRepo()

	userService, err := services.New(repo)

	if err != nil {
		fmt.Errorf("something went wrong")

		return
	}

	server := httpserver.NewServer(userService)

	server.Serve()
}
