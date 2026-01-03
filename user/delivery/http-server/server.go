package httpserver

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pooya/config"
	userhandler "github.com/pooya/delivery/http-server/user-handler"
	userservice "github.com/pooya/services"
)

type Server struct {
	userHandler userhandler.Handler
}

func NewServer(userService userservice.Service) Server {
	server := Server{
		userHandler: userhandler.New(userService),
	}

	return server
}

func (s Server) Serve() {
	cfg := config.Load("config.yml")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s.userHandler.SetUserRoute(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s", cfg.UserMicroservicePort.Port)))
}
