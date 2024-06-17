package handlers

import (
	"fmt"

	"github.com/Mayer-04/fiber-authentication/internal/models"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateRegisterData valida los datos del registro
func validateRegisterData(data *models.Register) error {
	if err := validate.Struct(data); err != nil {

		for _, err := range err.(validator.ValidationErrors) {

			message := fmt.Sprintf("the %s field is required", err.Field())
			return fmt.Errorf("the field entered is incorrect: %s", message)

		}
	}

	return nil
}

// ValidateLoginData valida los datos del inicio de sesión "login"
func validateLoginData(data *models.Login) error {
	if err := validate.Struct(data); err != nil {
		return fmt.Errorf("validate login: %w", err)
	}
	return nil
}
