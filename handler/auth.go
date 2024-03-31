package handler

import (
	"errors"
	"log"
	"time"

	"github.com/Mayer-04/fiber-authentication/database"
	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	var data models.Register
	db := database.DB

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Error parsing request body to a struct"})
	}

	if err := ValidateRegisterData(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	hash, err := HashPassword(data.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to hash password"})
	}

	newUser := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hash,
	}

	// Si hay un error al crear el usuario como conflicto de clave única "email" retornar un error
	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"success": false, "message": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "data": newUser})
}

func Login(c *fiber.Ctx) error {
	var data models.Login
	db := database.DB

	// Parsear el cuerpo de la solicitud
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Error parsing request body to a struct"})
	}

	// Validar datos de inicio de sesión
	if err := ValidateLoginData(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var user models.User
	// Buscar usuario por correo electrónico
	if err := db.Where("email = ?", data.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": "User not found"})
		}
		// Manejar otros errores de la base de datos
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Database error"})
	}

	// Comparar contraseña ingresada con la contraseña almacenada en la base de datos
	if !ComparePasswordAndHash(data.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// Generar token JWT
	token, err := GenerateToken(user)
	if err != nil {
		// Registrar el error
		log.Printf("failed to generate token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to generate token"})
	}

	// Creando una cookie
	cookie := &fiber.Cookie{
		Name:  "Authorization",
		Value: token,
		// Secure solo para HTTPS
		Secure: true,
		// HttpOnly solo puede ser accedida o leída por peticiones HTTP
		HTTPOnly: true,
		// SameSite controlar si la cookie puede ser compartida entre dominios "CORS"
		SameSite: fiber.CookieSameSiteNoneMode,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	// Agregar la cookie a la respuesta
	c.Cookie(cookie)

	// Retornar éxito y token JWT
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "token": token})
}
