package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func pong(route string) string {
	return fmt.Sprintf("%s\n", route)
}

func getBalance(address string) (string, error) {
	balance, err := token.BalanceOf(callOpts, common.HexToAddress(address))
	return balance.String(), err
}