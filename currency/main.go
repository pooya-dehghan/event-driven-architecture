package main

import (
	"fmt"
	"log"

	"github.com/expse/config"
	"github.com/expse/repository"
	"github.com/expse/services"
	"github.com/nats-io/nats.go"
)

type Server struct {
	natConn *nats.Conn
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	cfg := config.Load("config.yml")

	repo := repository.New(&cfg.MysqlConfig)

	svc := services.NewService(&repo)

	if err != nil {
		fmt.Println(err)
	}

	log.Println("nats connection was made")

	srv := &Server{
		natConn: nc,
	}

	srv.natConn.Subscribe("add-currency-request", func(msg *nats.Msg) {
		svc.CreateCurrencyRequest(msg)
	})

	select {}
}
