package database

import (
	"fmt"
	"log"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var err error

	config := config.LoadEnvVariables()
	dbURL := config.PostgresURL

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database")

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	DB = db

}
