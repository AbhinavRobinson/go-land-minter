package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()
	// Middlewares
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())
	// Routes
	setupRoutes(app)
	// Listen
	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		return
	}
}
