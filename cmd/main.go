package main

import (
	"fmt"
	"log"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/database"
	"github.com/Mayer-04/fiber-authentication/middlewares"
	"github.com/Mayer-04/fiber-authentication/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// Conexión a la base de datos
	database.ConnectDatabase()

	port := config.LoadEnvVariables().Port

	// Configuración de Fiber
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})

	// Configuración de CORS
	app.Use(cors.New(middlewares.GetCorsConfig()))

	// Configuración de rutas
	routes.SetupRoutes(app)

	address := fmt.Sprintf(":%d", port)

	log.Fatal(app.Listen(address))

}
