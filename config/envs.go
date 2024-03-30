package config

import (
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

// Variables para las diferentes variables de entorno
var (
	env              = ".env"
	Port             = "PORT"
	PostgresUser     = "POSTGRES_USER"
	PostgresPassword = "POSTGRES_PASSWORD"
	PostgresDB       = "POSTGRES_DB"
	PostgresURL      = "POSTGRES_URL"
	JwtSecret        = "JWT_SECRET"
)

func LoadEnvVariables() Envs {
	err := godotenv.Load(env)
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}
	// Parsear el puerto - 10 es el número maximo de digitos, 32 es el número maximo de bytes
	port, err := strconv.ParseUint(os.Getenv(Port), 10, 32)
	if err != nil {
		log.Fatal("failed to parse PORT environment variable")
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
