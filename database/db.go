package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB
var err error

func ConnectDatabase() {
	url := os.Getenv("DB_URL")

	if url == "" {
		log.Fatal("database credentials missing")
		os.Exit(1)
	}

	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})


	if err != nil {
		log.Fatal("Failed to connect to db")
		os.Exit(1)
	}

	fmt.Println("Connected To DB successfully")
}