package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

// UTILS

func ping(c *fiber.Ctx) error {
	return c.SendString(_pong(c.Route().Path))
}

// TOKEN

func getBalanceOfAddress(c *fiber.Ctx) error {
	res, err := _getBalance(c.Params("addr"))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString(res)
}

func getAllowanceForAddress(c *fiber.Ctx) error {
	res, err := _getAllowance(c.Params("from"), c.Params("to"))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString(res)
}

// LAND

func getLandOwner(c *fiber.Ctx) error {
	res, err := _getLandOwner(c.Params("id"))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString(res)
}

func getLandId(c *fiber.Ctx) error {
	res, err := _getLandId(c.Params("x"), c.Params("y"))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString(res)
}

func getLandCoordinates(c *fiber.Ctx) error {
	res, err := _getLandCoordinates(c.Params("id"))
	if err != nil {
		return c.SendString(err.Error())
	}
	resp, err := json.Marshal(res)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString(string(resp))
}

// SALE

func getSaleById(c *fiber.Ctx) error {
	res, err := _getSaleById(c.Params("id"))
	if err != nil {
		return c.SendString(err.Error())
	}
	resp, err := json.Marshal(res)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString(string(resp))
}
