package models

type Register struct {
	Name     string `json:"name" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gte=6"`
	Email    string `json:"email" validate:"required,email"`
}
