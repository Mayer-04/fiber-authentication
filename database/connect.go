package database

import (
	"log"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variable global que instancia la base de datos
var DB *gorm.DB

// ConnectDatabase establece la conexión con la base de datos y realiza la migración inicial
func ConnectDatabase() {

	config := config.LoadEnvVariables()
	dbURL := config.PostgresURL

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Realiza la migración inicial
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("connected to database")
}
