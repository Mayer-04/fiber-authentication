package handler

import (
	"fmt"
	"time"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hash password: %w", err)
	}
	return string(hashedBytes), nil
}

// CheckPasswordHash retorna un booleano.
// Si el error es nil, las contraseñas son iguales, retorna true.
// Si retorna un error, las contraseñas no son iguales, retorna false
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// * Funciones JWT

// Convertir la variable de entorno "JWT_SECRET" a []byte
var secretKey = []byte(config.LoadEnvVariables().JwtSecret)

func GenerateToken(user models.User) (string, error) {
	// Información que contiene datos sobre el usuario o entidad asociada con el token
	// Información que compartiremos con el cliente
	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		// Establece la expiración en 24 horas
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

// func GenerateToken(user models.User) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["id"] = user.ID
// 	claims["name"] = user.Name
// 	claims["username"] = user.UserName
// 	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

// 	tokenString, err := token.SignedString(secretKey)

// 	if err != nil {
// 		return "", fmt.Errorf("failed to generate token: %v", err)
// 	}

// 	return tokenString, nil
// }

// * Creación de la cookie

// CreateCookie Crea una nueva cookie que recibe como valor el token y otras configuraciones
func CreateCookie(token string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,                          // Solo para HTTPS
		HTTPOnly: true,                          // Solo puede ser accedida o leída por peticiones HTTP
		SameSite: fiber.CookieSameSiteNoneMode,  // Controla si la cookie puede ser compartida entre dominios "CORS"
		Expires:  time.Now().Add(3 * time.Hour), // Tiempo de expiración de la cookie - 3 horas
	}

}
