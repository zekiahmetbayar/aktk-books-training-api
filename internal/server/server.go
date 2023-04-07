package server

import (
	"fmt"
	"go-training/app/middleware"
	"go-training/app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start(port int) {
	// Create Fiber App
	app := fiber.New(fiber.Config{
		Prefork:      false,
		ErrorHandler: ErrorHandler,
	})

	// Initialize Logger
	app.Use(logger.New())

	// Middlewares
	app.Use(middleware.Fake)

	// Mount routes
	routes.Routes(app)

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
