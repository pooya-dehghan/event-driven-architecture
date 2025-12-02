package repository

import (
	"fmt"

	"github.com/expse/entity"
	"github.com/expse/params"
)

func (r *Repo) CreateCurrencyRequest(params params.CurrencyRequestParams) error {

	var currency = entity.CurrencyRequest{
		UserId:      params.UserID,
		Price:       params.Price,
		Description: params.Description,
	}

	if err := r.db.Create(&currency).Error; err != nil {
		fmt.Println("user creation had a problem", err)
	}

	return nil
}
