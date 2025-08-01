package main

import (
	"log"
	"natsas-place/config"
	"natsas-place/database"
	"natsas-place/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {

	cfg := config.GetConfig()

	database.ConnectBD()

	app := fiber.New(fiber.Config{
		AppName:               "Natsa's Place API",
		DisableStartupMessage: false,
		ErrorHandler:          customErrorHandler,
		BodyLimit:             10 * 1024 * 1024, // 10MB
	})

	// Global middlewares
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} - ${latency}\n",
	}))

	app.Use(recover.New())
	app.Use(requestid.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.FrontendURL,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
	}))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "Natsa's Place API",
			"version": "1.0.0",
		})
	})

	// Setup routes
	routes.Setup(app)

	// Start server
	log.Printf("🚀 Server starting on port %s", cfg.Port)
	log.Printf("🌍 Environment: %s", cfg.Environment)
	log.Printf("🔗 Frontend URL: %s", cfg.FrontendURL)

	log.Fatal(app.Listen(":" + cfg.Port))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"message": message,
		"code":    code,
	})
}
