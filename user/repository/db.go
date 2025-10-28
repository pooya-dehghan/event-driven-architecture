package repository

import (
	"fmt"
	"log"

	"github.com/pooya/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo() Repo {
	dsn := "root:gold552@tcp(127.0.0.1:3308)/gold?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("open db error: %v", err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)

	}

	fmt.Println("connection was made correctly")

	rp := Repo{db: db}

	return rp
}
