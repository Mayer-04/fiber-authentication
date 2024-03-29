package models

type Login struct {
	Email    string `json:"email" validate:"required,email,omitempty"`
	Password string `json:"password" validate:"required,gte=6,omitempty"`
}
