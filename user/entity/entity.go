package entity

import "time"

type User struct {
	Id          				string `json:"id" gorm:"column:id"`
	Name        				string `json:"name" gorm:"column:name"`
	PhoneNumber 				string `json:"phonenumber" gorm:"phone-number"`
	CreatedAt                   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt                   *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
