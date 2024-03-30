package config

import (
	"fmt"
	"time"

	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(LoadEnvVariables().JwtSecret) // Convertir a []byte

func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"name":     user.Name,
		"username": user.UserName,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
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
