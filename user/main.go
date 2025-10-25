package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"github.com/pooya/entity"
)

const PORT = "7366"

type Server struct {
	natConn *nats.Conn
}

func (nc *Server) AddExpenseHandler(c echo.Context) error{
	fmt.Println("adding")

	user := entity.User{
		Id:          "1",
		Name:        "pooya",
		PhoneNumber: "09202230930",
	}

	log.Println(user)

	if err := nc.natConn.Publish("add-expense", []byte("New Expense Added")); err != nil {
		log.Println("Error publishing to NATS:", err)
	}

	log.Printf("adding ends")

	return c.JSON(http.StatusOK, user)
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}

	defer nc.Drain()

	srv := &Server{natConn: nc}

	e := echo.New()
	e.POST("/add-expense", srv.AddExpenseHandler)

	e.Start(":" + PORT)
}
