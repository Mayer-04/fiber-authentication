package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Envs struct {
	Port             uint64
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresURL      string
	JwtSecret        string
}

func LoadEnvVariables() Envs {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}
	// Parse PORT - 10 is the maximum number of digits, 32 is the maximum number of bytes
	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 32)
	if err != nil {
		log.Fatal("failed to parse PORT environment variable")
	}

	return Envs{
		Port:             port,
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresURL:      os.Getenv("POSTGRES_URL"),
		JwtSecret:        os.Getenv("JWT_SECRET"),
	}
}
