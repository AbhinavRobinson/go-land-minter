package main

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *fiber.App) {
	// Add Groups
	_minter := app.Group("/minter")
	_token := _minter.Group("/token")
	_land := _minter.Group("/land")
	_sale := _minter.Group("/sale")

	// Token
	_token.Get("/balance/:addr", getBalanceOfAddress)
	_token.Get("/allowance/:from/:to", getAllowanceForAddress)

	// Land
	_land.Get("/owner/:id", getLandOwner)

	// Sale
	_sale.Get("/status/:id", getSaleById)
}