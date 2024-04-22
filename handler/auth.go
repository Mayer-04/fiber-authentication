package handler

import (
	"errors"
	"log"

	"github.com/Mayer-04/fiber-authentication/database"
	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	var data models.Register
	db := database.DB

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error parsing request body to a struct",
		})
	}

	if err := validateRegisterData(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	hash, err := hashPassword(data.Password)

	if err != nil {
		// Registrar si hay un error
		log.Printf("failed hash password %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to hash password"})
	}

	newUser := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hash,
	}

	// Si hay un error al crear el usuario como conflicto de clave única "email" retornar un error
	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "data": newUser})
}

func Login(c *fiber.Ctx) error {
	var data models.Login
	db := database.DB

	// Parsear el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error parsing request body to a struct",
		})
	}

	// Validar datos de inicio de sesión
	if err := validateLoginData(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var user models.User

	// Filtro para buscar el "primer" usuario que coincida con el correo electronico
	queryFilter := db.Where("email = ?", data.Email).First(&user)

	// Si tenemos un error al buscar el usuario por email
	if err := queryFilter.Error; err != nil {

		// Registrar el error
		log.Printf("queryFilter: %v", err)

		// Si el error es igual al error de registro no encontrado en GORM
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "User not found",
			})
		}
		// Manejar otros errores de la base de datos
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database error",
			"error":   err.Error(),
		})
	}

	// Comparar contraseña ingresada con la contraseña almacenada en la base de datos
	if !checkPasswordHash(data.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// Generar token JWT
	token, err := generateToken(user)
	if err != nil {
		// Registrar si hay un error
		log.Printf("failed to generate token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to generate token"})
	}

	cookie := createCookie(token)

	// Agregar la cookie a la respuesta
	c.Cookie(cookie)

	// Retornar éxito y token JWT
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}
