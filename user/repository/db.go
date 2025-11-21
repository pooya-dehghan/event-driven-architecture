package repository

import (
	"fmt"
	"log"

	"github.com/pooya/config"
	"github.com/pooya/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo struct {
	db     *gorm.DB
	config *config.MysqlDatabase
}

func NewRepo(config *config.MysqlDatabase) Repo {
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DatabasePass, config.Host, config.Port, config.DatabaseName)

	fmt.Println("dsn %v", dsn)

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
