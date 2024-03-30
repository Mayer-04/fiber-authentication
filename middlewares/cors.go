package middlewares

import "github.com/gofiber/fiber/v2/middleware/cors"

// GetCorsConfig retornar la configuración de CORS
func GetCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     "http://localhost:5173, http://localhost:3000, http://localhost:4321",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}
}