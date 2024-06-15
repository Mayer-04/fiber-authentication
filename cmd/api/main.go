package main

import (
	"fmt"
	"log"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/internal/database"
	"github.com/Mayer-04/fiber-authentication/internal/server"
)

func main() {

	// Conexión a la base de datos
	database.ConnectDatabase()

	port := config.LoadEnvVariables().Port

	// Configuración de Fiber
	app := server.New()

	// Configuración de rutas
	app.SetupRoutes()

	address := fmt.Sprintf(":%d", port)

	log.Fatal(app.Listen(address))

}
