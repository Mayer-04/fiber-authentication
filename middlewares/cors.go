package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Métodos HTTP que acepta la configuración de cors
var methods = strings.Join([]string{
	fiber.MethodGet,
	fiber.MethodPost,
}, ",")

// GetCorsConfig retornar la configuración de CORS
func GetCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     "http://localhost:5173, http://localhost:3000, http://localhost:4321",
		AllowMethods:     methods,
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}
}
