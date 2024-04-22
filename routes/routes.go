package routes

import (
	"github.com/Mayer-04/fiber-authentication/handler"
	"github.com/Mayer-04/fiber-authentication/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/v1")
	auth := api.Group("/auth")

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	api.Get("/users", middlewares.VerifyToken, handler.FindAllUsers)
}
