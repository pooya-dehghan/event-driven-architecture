package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	uri := "mongodb://root:gold552@localhost:27017/gold?authSource=admin"

	// Set client options
	clientOpts := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	// Access the database and collection
	db := client.Database("gold")
	usersCollection := db.Collection("users")

	//usersCollection.

	srv.natConn.Subscribe("add-expense", func(msg *nats.Msg) {
		log.Println("hi i was called")
		log.Println("message: ", msg)
	})

	select {}
}
