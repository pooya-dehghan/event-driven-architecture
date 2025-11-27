package repository

import (
	"fmt"
	"log"

	"github.com/expse/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo struct {
	db     *gorm.DB
	config *config.MysqlConfig
}

func New(config *config.MysqlConfig) Repo {
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DatabasePass, config.Host, config.Port, config.DatabaseName)

	fmt.Println("dsn %v", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("open db error: %v", err)
	}

	return Repo{
		db:     db,
		config: config,
	}
}
