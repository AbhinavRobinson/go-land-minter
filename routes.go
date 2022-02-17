package main

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *fiber.App) {
	// Add Groups
	minter := app.Group("/minter")  // /minter
	token := minter.Group("/token") // /minter/token
	land := minter.Group("/land")   // /minter/land
	sale := minter.Group("/sale")   // /minter/sale

	// Add Ping routes
	minter.Get("/", ping) 	// pongs /minter
	token.Get("/", ping)  	// pongs /minter/token
	land.Get("/", ping)   	// pongs /minter/land
	sale.Get("/", ping)   	// pongs /minter/sale
}

func ping(c *fiber.Ctx) error {
	return c.SendString(pong(c.Route().Path))
}
