package repository

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo struct { 
	db *gorm.DB
}

func NewRepo (database *gorm.DB) *Repo{
	dsn := "root:gold552@tcp(127.0.0.1:3308)/gold?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if(err != nil){
		log.Printf("open db error: %v", err)
	}

	rp := &Repo{db: db};

	return rp
}

