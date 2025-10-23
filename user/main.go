package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
	"github.com/pooya/entity"
)

const PORT = "7366"

type Server struct {
	natConn *nats.Conn
}

func EchoHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("something in here")

	log.Println("handler")

}

func (nc *Server) AddExpenseHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

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
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}

	defer nc.Drain()

	srv := &Server{natConn: nc}

	http.HandleFunc("/", EchoHandler)

	http.HandleFunc("/add-expense", srv.AddExpenseHandler)

	http.ListenAndServe(":"+PORT, nil)
}
