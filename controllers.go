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

func getAllowance(owner string, spender string) (string, error) {
	allowance, err := token.Allowance(callOpts, common.HexToAddress(owner), common.HexToAddress(spender))
	return allowance.String(), err
}

// LAND

// SALE