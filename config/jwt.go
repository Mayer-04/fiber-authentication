package config

import (
	"fmt"
	"time"

	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(LoadEnvVariables().JwtSecret) // Convertir la variable de entorno a []byte

func GenerateToken(user models.User) (string, error) {
	// Informacion que contiene datos sobre el usuario o entidad asociada con el token
	claims := jwt.MapClaims{
		"id":       user.ID,
		"name":     user.Name,
		"username": user.UserName,
		// Establece la expiración en 24 horas
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	// Crea un nuevo token con el metodo HS256 y establece las claims del token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Genera y firma el token - La firma espera un []byte para la clave secreta y devuelve un string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return tokenString, nil
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
