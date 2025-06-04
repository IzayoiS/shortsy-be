package database

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"shortsy/models"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_URL")
	db,err :=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect database")
	}

	DB=db

	db.AutoMigrate(&models.Shortsy{})
}