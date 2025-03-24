package config

import (
	"fmt"
	"log"
	// "os"

	"github.com/carlkiptoo/backend/models"
	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "host=localhost port=5432 user=postgres password=Carlos dbname=taskmanager sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.AutoMigrate(&models.User{})


	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	DB = db

	fmt.Println("Database connected successfully")

}