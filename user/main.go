package main

import (
	"fmt"

	httpserver "github.com/pooya/delivery/http-server"
	server "github.com/pooya/delivery/rabbit-mq"
	"github.com/pooya/repository"
	userservice "github.com/pooya/services"
)

func main() {

	repo := repository.NewRepo()

	rabbitConn := server.NewRabbitMQConnection()

	userService, err := userservice.New(repo, rabbitConn)

	if err != nil {
		fmt.Errorf("something went wrong")

		return
	}

	server := httpserver.NewServer(userService)

	server.Serve()
}
