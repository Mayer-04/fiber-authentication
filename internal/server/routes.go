package server

import (
	"github.com/Mayer-04/fiber-authentication/internal/handlers"
	"github.com/Mayer-04/fiber-authentication/internal/middlewares"
)

// SetupRoutes configura las rutas de la API
func (s *FiberServer) SetupRoutes() {
	api := s.Group("/v1")
	authRoutes := api.Group("/auth")

	authRoutes.Post("/register", handlers.Register)
	authRoutes.Post("/login", handlers.Login)
	authRoutes.Post("/logout", handlers.Logout)
	api.Get("/users", middlewares.VerifyToken, handlers.FindAllUsers)
}
