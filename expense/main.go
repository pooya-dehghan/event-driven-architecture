package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/expse/entity"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	natConn *nats.Conn
}

func ConnectingToMongoDB() (*mongo.Client, error) {
	uri := "mongodb://root:gold552@localhost:27019/gold?authSource=admin"

	clientOpts := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("gold")
	expenseColl := db.Collection("expense")

	expense := entity.Expense{
		UserId: "1",
		Desc:   "this expense was to get a pizza",
		Amount: 10000,
	}

	res, err := expenseColl.InsertOne(ctx, expense)

	if err != nil {
		fmt.Println("could not insert expense")
	}

	fmt.Println("res***: ", res)

	return client, nil
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	if err != nil {
		fmt.Println(err)
	}

	ConnectingToMongoDB()
	log.Println("nats connection was made")

	srv := &Server{
		natConn: nc,
	}

	srv.natConn.Subscribe("add-expense", func(msg *nats.Msg) {
		log.Println("hi i was called")
		log.Println("message: ", msg)
	})

	select {}
}
