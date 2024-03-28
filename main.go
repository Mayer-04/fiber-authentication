package main

import (
	"fmt"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	port := config.LoadEnvVariables().Port

	database.ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	addr := fmt.Sprintf(":%d", port)

	app.Listen(addr)

}
