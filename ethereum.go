package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func setupContracts() {
	provider := os.Getenv("PROVIDER")
	client, errClient := ethclient.Dial(provider)
	if errClient != nil {
		log.Fatal(errClient)
	}
	address := common.HexToAddress("0x4E0732efCdF0Cf92C48439535CD763de06FE353a")
	instance, errToken := NewToken(address, client)
	if errToken != nil {
		log.Fatal(errToken)
	}
	fmt.Println("🚀 Contracts are loaded!")
	_ = instance
}
