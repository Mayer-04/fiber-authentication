package database

import (
	"fmt"
	"log"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDatabase() {
	config := config.LoadEnvVariables()
	dbURL := config.PostgresURL

	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database")

	err = database.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	DB = Dbinstance{Db: database}

}
