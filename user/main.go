package main

import (
	"github.com/pooya/config"
	httpserver "github.com/pooya/delivery/http-server"
	server "github.com/pooya/delivery/rabbit-mq"
	"github.com/pooya/repository"
	userservice "github.com/pooya/services"
)

func main() {
	cfg := config.Load("config.yml")

	repo := repository.NewRepo(&cfg.MysqlDatabase)

	rabbitConn := server.NewRabbitMQConnection()

	userService, err := userservice.New(repo, rabbitConn)

	server := httpserver.NewServer(userService, cfg)

	server.Serve()

	if err != nil {
		panic("something happend on init")
	}
}
