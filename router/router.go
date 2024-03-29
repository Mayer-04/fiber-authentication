package router

import (
	"github.com/Mayer-04/fiber-authentication/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	routes := app.Group("/api", logger.New())

	authRoutes := routes.Group("/auth")
	authRoutes.Post("/register", handler.Register)
	authRoutes.Post("/login", handler.Login)
}
