package server

import (
	"github.com/Mayer-04/fiber-authentication/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	*fiber.App
}

// New crea una nueva instancia de FiberServer
func New() *FiberServer {

	config := fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	}

	// Instancia del servidor
	server := &FiberServer{
		App: fiber.New(config),
	}

	// Configuración de CORS
	server.Use(cors.New(middlewares.GetCorsConfig()))

	return server

}
