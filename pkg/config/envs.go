package config

import (
	"fmt"
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

const env = ".env"

// LoadEnvVariables carga las variables de entorno y retorna un objeto de tipo Envs con los valores.
func LoadEnvVariables() Envs {

	if err := godotenv.Load(env); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	portStr := getEnvValue(Port)
	port, err := parsePort(portStr)

	if err != nil {
		log.Fatal(err)
	}

	return Envs{
		Port:             port,
		PostgresUser:     getEnvValue(PostgresUser),
		PostgresPassword: getEnvValue(PostgresPassword),
		PostgresDB:       getEnvValue(PostgresDB),
		PostgresURL:      getEnvValue(PostgresURL),
		JwtSecret:        getEnvValue(JwtSecret),
	}
}

// getEnvValue recupera el valor de la variable de entorno especificada.
// Registra un error y sale si la variable de entorno no está configurada.
func getEnvValue(env string) string {
	value, ok := os.LookupEnv(env)
	if !ok {
		log.Fatalf("environment variable %q not set", env)
	}
	return value
}

// parsePort parsea el puerto de la variable de entorno a un valor de tipo uint64.
func parsePort(portStr string) (uint64, error) {

	if portStr == "" {
		return 8080, nil
	}

	// Parsear el puerto - 10 es el número maximo de digitos, 32 es el número maximo de bytes
	port, err := strconv.ParseUint(portStr, 10, 32)

	if err != nil {
		return 0, fmt.Errorf("failed to parse PORT environment variable: %w", err)
	}

	return port, nil
}
