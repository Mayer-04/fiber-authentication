package handler

import (
	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateRegisterData valida los datos del registro
func ValidateRegisterData(data *models.Register) error {
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}

// ValidateLoginData valida los datos del inicio de sesión "login"
func ValidateLoginData(data *models.Login) error {
	if err := validate.Struct(data); err != nil {
		return err
	}

	return nil
}
