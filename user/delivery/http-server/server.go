package httpserver

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
	"github.com/pooya/config"
	userhandler "github.com/pooya/delivery/http-server/user-handler"
	userservice "github.com/pooya/services"
)

type Server struct {
	userHandler userhandler.Handler
	config      *config.UserConfig
}

func NewServer(userService userservice.Service, cfg *config.UserConfig) Server {
	server := Server{
		userHandler: userhandler.New(userService),
		config:      cfg,
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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s.userHandler.SetUserRoute(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s", s.config.UserMicroservicePort.Port)))
}
