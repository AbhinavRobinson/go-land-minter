package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// UTILS

func _pong(route string) string {
	return fmt.Sprintf("%s\n", route)
}

// TOKEN

func _getBalance(address string) (string, error) {
	balance, err := token.BalanceOf(callOpts, common.HexToAddress(address))
	return balance.String(), err
}

func _getAllowance(owner string, spender string) (string, error) {
	allowance, err := token.Allowance(callOpts, common.HexToAddress(owner), common.HexToAddress(spender))
	return allowance.String(), err
}

// LAND

func _getLandOwner(landIdString string) (string, error) {
	landId := new(big.Int)
	landId, ok := landId.SetString(landIdString, 10)
	if !ok {
		return "", fmt.Errorf("invalid landId")
	}
	owner, err := land.OwnerOf(callOpts, landId)
	return owner.String(), err
}

func _getLandId(xCoordinate string, yCoordinate string) (string, error) {
	x, ok := new(big.Int).SetString(xCoordinate, 10)
	if !ok {
		return "", fmt.Errorf("invalid xCoordinate")
	}
	y, ok := new(big.Int).SetString(yCoordinate, 10)
	if !ok {
		return "", fmt.Errorf("invalid yCoordinate")
	}
	landId, err := land.GetIdFromCoordinates(callOpts, x, y)
	return landId.String(), err
}

func _getLandCoordinates(landIdString string) (Coordinates, error) {
	landId := new(big.Int)
	landId, ok := landId.SetString(landIdString, 10)
	if !ok {
		return Coordinates{nil,nil}, fmt.Errorf("invalid landId")
	}
	x, y, err := land.GetCoordinatesFromId(callOpts, landId)
	return Coordinates{x,y}, err
}

// SALE

func _getSaleById(saleIdString string) (Sale, error) {
	saleId := new(big.Int)
	saleId, ok := saleId.SetString(saleIdString, 10)
	if !ok {
		return Sale{false, nil}, fmt.Errorf("invalid saleId")
	}
	res, err := sale.GetSale(callOpts, saleId)
	return Sale{res.IsActive, res.Price}, err
}
