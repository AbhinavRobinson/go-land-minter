package main

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *fiber.App) {
	// Add Groups
	_minter := app.Group("/minter")   // /minter
	_token := _minter.Group("/token") // /minter/token
	_land := _minter.Group("/land")   // /minter/land
	_sale := _minter.Group("/sale")   // /minter/sale

	// Add Ping routes
	_minter.Get("/", ping) // pongs /minter
	_token.Get("/", ping)  // pongs /minter/token
	_land.Get("/", ping)   // pongs /minter/land
	_sale.Get("/", ping)   // pongs /minter/sale

	// Token
	_token.Get("/:address", getAddress)
}

func ping(c *fiber.Ctx) error {
	return c.SendString(pong(c.Route().Path))
}

func getAddress(c *fiber.Ctx) error {
	res, err := getBalance(c.Params("address"))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString(res)
}
