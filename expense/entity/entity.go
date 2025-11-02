package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	Id        primitive.ObjectID `bson:"id"`
	UserId    string             `bson:"userid"`
	Amount    int                `bson:"amount"`
	Desc      string             `bson:"desc"`
	CreatedAt time.Time          `bson:"created_at`
	UpdatedAt time.Time          `bson:"updated_at"`
}
