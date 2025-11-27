package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

type Server struct {
	natConn *nats.Conn
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	if err != nil {
		fmt.Println(err)
	}

	log.Println("nats connection was made")

	srv := &Server{
		natConn: nc,
	}

	srv.natConn.Subscribe("add-currency-request", func(msg *nats.Msg) {
		log.Println("message: ", msg)
	})

	select {}
}
