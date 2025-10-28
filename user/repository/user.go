package repository

import (
	"fmt"

	"github.com/pooya/entity"
)

func (r *Repo) AddExpense(amount int) error {
	user := entity.User{
		Name:        "pooya",
		PhoneNumber: "09202230930",
	}

	if err := r.db.Create(&user).Error; err != nil {
		fmt.Println("user creation had a problem", err)
	}

	return nil
}
