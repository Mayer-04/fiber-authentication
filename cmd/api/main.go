package main

import (
	"fmt"
	"log"

	"github.com/Mayer-04/fiber-authentication/config"
	"github.com/Mayer-04/fiber-authentication/internal/database"
	"github.com/Mayer-04/fiber-authentication/internal/server"
)

func main() {

	// Carga de variables de entorno
	env := config.LoadEnvVariables()

	// Conexión a la base de datos
	database.ConnectDatabase(env.PostgresURL)

	// Configuración de Fiber
	app := server.New()

	// Configuración de rutas
	app.SetupRoutes()

	address := fmt.Sprintf(":%d", env.Port)

	log.Fatal(app.Listen(address))

}
