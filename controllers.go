package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// UTILS

func pong(route string) string {
	return fmt.Sprintf("%s\n", route)
}

// TOKEN

func getBalance(address string) (string, error) {
	balance, err := token.BalanceOf(callOpts, common.HexToAddress(address))
	return balance.String(), err
}

// LAND

// SALE