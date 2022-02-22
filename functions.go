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
