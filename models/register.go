package models

type Register struct {
	Name     string `json:"name" validate:"required,omitempty"`
	Password string `json:"password" validate:"required,gte=6,omitempty"`
	Email    string `json:"email" validate:"required,email"`
}
