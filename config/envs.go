package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Estructura para las variables de entorno
type Envs struct {
	Port             uint64
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresURL      string
	JwtSecret        string
}

// Constantes para las diferentes variables de entorno
const (
	Port             = "PORT"
	PostgresUser     = "POSTGRES_USER"
	PostgresPassword = "POSTGRES_PASSWORD"
	PostgresDB       = "POSTGRES_DB"
	PostgresURL      = "POSTGRES_URL"
	JwtSecret        = "JWT_SECRET"
)

func LoadEnvVariables() Envs {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	portStr := os.Getenv(Port)
	port, err := parsePort(portStr)

	if err != nil {
		log.Fatal(err)
	}

	return Envs{
		Port:             port,
		PostgresUser:     os.Getenv(PostgresUser),
		PostgresPassword: os.Getenv(PostgresPassword),
		PostgresDB:       os.Getenv(PostgresDB),
		PostgresURL:      os.Getenv(PostgresURL),
		JwtSecret:        os.Getenv(JwtSecret),
	}
}

func parsePort(portStr string) (uint64, error) {

	if portStr == "" {
		return 8080, nil
	}

	// Parsear el puerto - 10 es el número maximo de digitos, 32 es el número maximo de bytes
	port, err := strconv.ParseUint(portStr, 10, 32)

	if err != nil {
		return 0, errors.New("failed to parse PORT environment variable")
	}

	return port, nil
}
