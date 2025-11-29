package entity

import (
	"gorm.io/gorm"
)

type CurrencyRequest struct {
	gorm.Model
	Id          uint         `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId      uint         `json:"userId" gorm:"column:userId"`
	Description string       `json:"description" gorm:"column:description"`
	Price       uint         `json:"price" gorm:"column:price"`
	Type        CurrencyType `json:"type" gorm:"type:VARCHAR(20)"`
}

type CurrencyType string

const (
	Gold   CurrencyType = "Gold"
	Silver CurrencyType = "Silver"
	Oil    CurrencyType = "Oil"
)
