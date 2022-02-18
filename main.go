package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"

	"os"
)

func main() {
	// Load ENV
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	// Load Contracts
	setupContracts()
	// Load Fiber
	app := fiber.New()
	// Middlewares
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())
	// Routes
	setupRoutes(app)
	// Listen
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	_ = app.Listen(fmt.Sprintf("%s:%s", host, port))
}
