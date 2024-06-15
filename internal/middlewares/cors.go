package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// corsOrigins es un slice que contiene los origenes aceptados por la configuración de CORS
var corsOrigins = []string{
	"http://localhost:5173",
	"http://localhost:3000",
	"http://localhost:4321",
	"http://localhost:5000",
}

// origins es una cadena que contiene la lista de origenes aceptados separados por comas
var origins = strings.Join(corsOrigins, ",")

// fiberMethods es un slice que contiene los métodos HTTP aceptados por la configuración de CORS
var fiberMethods = []string{
	fiber.MethodGet,
	fiber.MethodPost,
}

// methods es una cadena que contiene la lista de métodos HTTP aceptados separados por comas
var methods = strings.Join(fiberMethods, ",")

// GetCorsConfig retorna la configuración de CORS predeterminada para la aplicación
func GetCorsConfig() cors.Config {
	return cors.Config{
		// AllowOrigins define los orígenes permitidos para las solicitudes CORS
		AllowOrigins: origins,
		// AllowMethods define los métodos HTTP permitidos en las solicitudes CORS
		AllowMethods: methods,
		// AllowCredentials indica si se permiten credenciales en las solicitudes CORS
		AllowCredentials: true,
		// AllowHeaders define los encabezados HTTP permitidos en las solicitudes CORS
		AllowHeaders: "Origin, Content-Type, Accept",
	}
}
