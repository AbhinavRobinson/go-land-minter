package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var token *Token
var cl *ethclient.Client

var ctx = context.Background()
var callOpts = &bind.CallOpts{Context: ctx, Pending: false}

func setupContracts() {
	provider := os.Getenv("PROVIDER")
	client, errClient := ethclient.Dial(provider)
	if errClient != nil {
		log.Fatal(errClient)
	}
	cl = client
	address := common.HexToAddress("4E0732efCdF0Cf92C48439535CD763de06FE353a")
	t, errToken := NewToken(address, client)
	if errToken != nil {
		log.Fatal(errToken)
	}
	token = t
	tokenName, err := token.Name(callOpts)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	fmt.Printf("Token Loaded: %s\n", tokenName)
	fmt.Println("🚀 Contracts are loaded!")
}
