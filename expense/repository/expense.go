package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/expse/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	mongoDb *mongo.Database
}

func (r *Repository) CreateExpense() (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	expColl := r.mongoDb.Collection("Expense")

	expense := entity.Expense{
		UserId: "1",
		Desc:   "this expense was to get a pizza",
		Amount: 10000,
	}

	exp, err := expColl.InsertOne(ctx, expense)

	if err != nil {
		fmt.Println("error on insertion")
	}

	return exp, nil
}
