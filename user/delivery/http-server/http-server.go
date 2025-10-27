package httpserver

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"github.com/pooya/services"
)

const PORT = "7366"

type Server struct {
	userService services.UserService
}

func (s Server) NewServer(userSrv services.UserService) Server {

	server := Server{
		userService: userSrv,
	}

	return server
}

func (s Server) Serve() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}

	defer nc.Drain()

	e := echo.New()
	e.POST("/add-expense", s.addExpense)

	e.Start(":" + PORT)
}
