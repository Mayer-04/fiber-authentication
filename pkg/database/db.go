package database

import (
	"fmt"
	"log"

	"github.com/Mayer-04/fiber-authentication/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	config := config.LoadEnvVariables()
	dbURL := config.PostgresURL

	DB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database")

	return DB

}
