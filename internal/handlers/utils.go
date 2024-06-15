package handlers

import (
	"fmt"
	"time"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// CookieName es el nombre de la cookie que contiene el token JWT.
const CookieName = "Authorization"

// hashPassword toma una contraseña y devuelve una nueva contraseña hasheada con bcrypt.
func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hash password: %w", err)
	}
	return string(hashedBytes), nil
}

// CheckPasswordHash retorna un booleano.
//
// Si el error es nil, las contraseñas son iguales, retorna true.
//
// Si retorna un error, las contraseñas no son iguales, retorna false.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// secretKey contiene la clave secreta para firmar y verificar tokens JWT.
var secretKey = []byte(config.LoadEnvVariables().JwtSecret)

func generateToken(user models.User) (string, error) {
	// Información que contiene datos sobre el usuario o entidad asociada con el token
	// Información que compartiremos con el cliente
	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		// Establece la expiración del token en 24 horas desde ahora
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	// Crea un nuevo token con el metodo HS256 y establece las claims del token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Genera y firma el token - La firma espera un []byte para la clave secreta y devuelve un string
	jwtString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return jwtString, nil
}

// VerifyToken verifica la validez de un token JWT dado.
func VerifyToken(tokenString string) error {
	// jwt.Parse analiza y verifica la validez de un token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return fmt.Errorf("parse token: %w", err)
	}

	// Valid especifica si el token es válido. Se completa cuando se analiza/verifica el token "jwt.Parse"
	if !token.Valid {
		return fmt.Errorf("invalid token: %w", err)
	}

	return nil
}

/*
func GenerateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["username"] = user.UserName
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}
*/

// createCookie crea una nueva cookie que contiene el token JWT.
func createCookie(token string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     CookieName,
		Value:    token,
		Secure:   true,                          // Solo para HTTPS
		HTTPOnly: true,                          // Solo accesible por peticiones HTTP
		SameSite: fiber.CookieSameSiteNoneMode,  // Controla la política de CORS
		Expires:  time.Now().Add(3 * time.Hour), // Expira en 3 horas
	}
}

// deleteCookie elimina la cookie de autorización.
//
// Devuelve un nuevo objeto *fiber.Cookie con el mismo nombre pero con una fecha de expiración en el pasado.
func deleteCookie() *fiber.Cookie {
	// Crear una cookie con el mismo nombre pero con una fecha de expiración en el pasado
	expiredCookie := new(fiber.Cookie)

	// Nombre de la cookie
	expiredCookie.Name = CookieName

	// Valor de la cookie se establece en una cadena vacía
	expiredCookie.Value = ""

	// Solo accesible por peticiones HTTP
	expiredCookie.HTTPOnly = true

	// Fecha de expiración en el pasado
	expiredCookie.Expires = time.Now().Add(-24 * time.Hour)

	return expiredCookie
}
