package routes

import (
	"github.com/Mayer-04/fiber-authentication/handler"
	"github.com/Mayer-04/fiber-authentication/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	auth := api.Group("/auth")

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Get("/user", middlewares.VerifyToken, func(c *fiber.Ctx) error {
		return c.SendString("Ruta privada")
	})
}
