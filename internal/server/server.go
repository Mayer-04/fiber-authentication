package server

import (
	"github.com/Mayer-04/fiber-authentication/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	app := &FiberServer{
		App: fiber.New(config),
	}

	// Configuración de Middlewares
	setupMiddlewares(app)

	return app

}

// SetupMiddlewares configura los Middlewares para la aplicación
func setupMiddlewares(app *FiberServer) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New(middlewares.GetCORSConfig()))
}
