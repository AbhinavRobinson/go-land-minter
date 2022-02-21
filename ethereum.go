package main

import (
	"context"
	"fmt"
	TOKEN "go_minter/contracts/token"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var token *TOKEN.Token
var cl *ethclient.Client

var ctx = context.Background()
var callOpts = &bind.CallOpts{Context: ctx, Pending: false}

func setupContracts() {
	provider := os.Getenv("PROVIDER")
	address := common.HexToAddress("4E0732efCdF0Cf92C48439535CD763de06FE353a")

	client, errClient := ethclient.Dial(provider)
	if errClient != nil {
		log.Fatalf("Error connecting to client: %s", errClient)
	}
	cl = client

	t, errToken := TOKEN.NewToken(address, client)
	if errToken != nil {
		log.Fatalf("Some error occurred. Err: %s", errToken)
	}
	token = t

	tokenName, err := token.Name(callOpts)
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}
	fmt.Printf("Token Loaded: %s\n", tokenName)
	fmt.Println("🚀 Contracts are loaded!")
}
