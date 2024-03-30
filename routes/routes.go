package routes

import (
	"github.com/Mayer-04/fiber-authentication/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	auth := api.Group("/auth")

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
}
