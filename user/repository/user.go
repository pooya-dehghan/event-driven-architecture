package repository

import (
	"context"

	"github.com/pooya/entity"
	"gorm.io/gorm"
)

func (r *Repo) AddExpense(amount int) error {
	user := entity.User{
		Name:        "pooya",
		PhoneNumber: "09202230930",
	}

	ctx := context.Background()
	err := gorm.G[entity.User](r.db).Create(ctx, &user)

	if err != nil {
		return err
	}
	return nil
}
