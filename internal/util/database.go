package util

import (
	"log"
	"tasko/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBCon *gorm.DB

func ConnectDatabase() {
	var err error
	DBCon, err = gorm.Open(postgres.Open("user=postgres password=postgres dbname=postgres sslmode=disable"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DBCon.AutoMigrate(&model.Task{})
}
