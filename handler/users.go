package handler

import (
	"github.com/Mayer-04/fiber-authentication/database"
	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/gofiber/fiber/v2"
)

// FindAllUsers recupera todos los usuarios de la base de datos.
func FindAllUsers(c *fiber.Ctx) error {
	var users []models.User

	db := database.DB

	// Método "Find" de GORM para recuperar todos los usuarios
	// Si ocurrió un error, devuelve un estado 404 y una respuesta JSON
	if result := db.Find(&users); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find users",
			"error":   result.Error,
		})
	}

	// Si no ocurrió ningún error, devuelve un estado 200 y una respuesta JSON con los usuarios
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}
