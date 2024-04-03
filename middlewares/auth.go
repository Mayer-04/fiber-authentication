package middlewares

import (
	"log"

	"github.com/Mayer-04/fiber-authentication/handler"
	"github.com/gofiber/fiber/v2"
)

func VerifyToken(c *fiber.Ctx) error {
	cookie := c.Cookies("Authorization")

	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Unauthorized, no token provided"})
	}

	err := handler.VerifyToken(cookie)

	if err != nil {
		log.Printf("failed to generate token: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Invalid Token"})
	}

	return c.Next()
}
