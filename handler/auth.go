package handler

import (
	"errors"
	"log"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/database"
	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

	// Si hay un error al crear el usuario como conflicto de clave única "email" retornar un error
	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "data": newUser})
}

func Login(c *fiber.Ctx) error {
	var data models.Login
	db := database.DB

	// Parsear el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Error parsing request"})
	}

	// Validar datos de inicio de sesión
	if err := ValidateLoginData(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	// Buscar usuario por correo electrónico
	var user models.User
	if err := db.Where("email = ?", data.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": "User not found"})
		}
		// Manejar otros errores de la base de datos
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Database error"})
	}

	// Comparar contraseña ingresada con la contraseña almacenada
	if !ComparePasswordAndHash(data.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Invalid credentials"})
	}

	// Generar token JWT
	token, err := config.GenerateToken(user)
	if err != nil {
		// Registrar el error
		log.Printf("Failed to generate token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to generate token"})
	}

	// Retornar éxito y token JWT
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Success login", "token": token})
}
