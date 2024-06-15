package middlewares

import (
	"log"

	"github.com/Mayer-04/fiber-authentication/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

// VerifyToken verifica la validez de un token JWT dado.
func VerifyToken(c *fiber.Ctx) error {

	token := c.Cookies(handlers.CookieName)

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized, no token provided",
		})
	}

	if err := handlers.VerifyToken(token); err != nil {
		log.Printf("failed to generate token: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Token",
		})
	}

	return c.Next()
}
