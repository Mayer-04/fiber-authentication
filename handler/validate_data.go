package handler

import (
	"github.com/Mayer-04/fiber-authentication/models"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRegisterData(formData *models.Register) error {
	if err := validate.Struct(formData); err != nil {
		return err
	}
	return nil
}

func ValidateLoginData(formData *models.Login) error {
	if err := validate.Struct(formData); err != nil {
		return err
	}

	return nil
}
