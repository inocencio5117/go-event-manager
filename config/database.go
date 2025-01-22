package config

import (
	"fmt"
	"log"
	"os"

	"github.com/inocencio5117/go-event-manager/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	mySqlDatabase := os.Getenv("MYSQL_DATABASE")
	mySqlUser := os.Getenv("MYSQL_USER")
	mySqlPassword := os.Getenv("MYSQL_PASSWORD")
	mySqlHost := os.Getenv("MYSQL_HOST")
	mySqlPort := os.Getenv("MYSQL_PORT")

	dsn :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mySqlUser, mySqlPassword, mySqlHost, mySqlPort, mySqlDatabase,
		)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database

	log.Println("Successfully connected to the database")

	AutoMigrate(DB)
}

func AutoMigrate(connection *gorm.DB) {
	err := connection.Debug().AutoMigrate(&models.Event{}, &models.Attendee{})

	if err != nil {
		log.Fatal("Failed to migrate tables:", err)
		return
	}

	log.Println("Successfully migrated tables")
}
