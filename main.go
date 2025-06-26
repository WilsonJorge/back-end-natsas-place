package main

import (
	"log"
	"natsas-place/config"
	"natsas-place/database"
	"natsas-place/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// loading environment variables
	config.GetConfig()

	// connect to database
	database.ConnectBD()

	// initialize fiber
	app := fiber.New(fiber.Config{
		AppName: "Natsa place",
	})

	// middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup das rotas
	routes.Setup(app)

	// Inicia o servidor
	log.Fatal(app.Listen(":3001"))
}
