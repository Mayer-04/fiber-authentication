package handler

import (
	"fmt"
	"time"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// ComparePasswordAndHash retorna un booleano.
// Si el error es nil, las contraseñas son iguales, retorna true.
// Si retorna un error, las contraseñas no son iguales, retorna false
func ComparePasswordAndHash(password, hash string) bool {
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
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return jwtString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
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
