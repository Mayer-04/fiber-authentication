package handler

import (
	"github.com/Mayer-04/fiber-authentication/database"
	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var formData models.Register
	db := database.DB.Db

	if err := c.BodyParser(&formData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Invalid request"})
	}

	if err := ValidateRegisterData(&formData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	newUser := models.User{
		Name:     formData.Name,
		Email:    formData.Email,
		Password: formData.Password,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "data": &newUser})
}

func Login(c *fiber.Ctx) error {

	return c.SendString("Login")
}
