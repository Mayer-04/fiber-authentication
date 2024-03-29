package handler

import (
	"github.com/Mayer-04/fiber-authentication/database"
	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data models.Register
	db := database.DB

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Invalid request"})
	}

	if err := ValidateRegisterData(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	hash, err := HashPassword(data.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to hash password"})
	}

	newUser := models.User{
		Name:     data.Name,
		UserName: data.UserName,
		Email:    data.Email,
		Password: hash,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "data": newUser})
}

func Login(c *fiber.Ctx) error {

	return c.SendString("Login")
}
