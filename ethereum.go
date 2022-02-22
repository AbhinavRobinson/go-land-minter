package main

import (
	"context"
	"fmt"
	LAND "go_minter/contracts/land"
	SALE "go_minter/contracts/sale"
	TOKEN "go_minter/contracts/token"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var token *TOKEN.Token
var land *LAND.Main
var sale *SALE.Main

var ctx = context.Background()
var callOpts = &bind.CallOpts{Context: ctx, Pending: false}

func setupContracts() {
	provider := os.Getenv("PROVIDER")
	tokenAddress := common.HexToAddress("4E0732efCdF0Cf92C48439535CD763de06FE353a")
	landAddress := common.HexToAddress("5f1994949073dA9252D79ee53e3C91255Dc90070")
	saleAddress := common.HexToAddress("751d2375C2Db72Af22350Db58c3F44a89F03e08c")
	var err error

	// INIT ETH CLIENT

	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatalf("Error connecting to client: %s", err)
	}

	// LOAD CONTRACTS

	t, err := TOKEN.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatalf("Some error occurred in TOKEN. Err: %s", err)
	}
	token = t

	l, err := LAND.NewMain(landAddress, client)
	if err != nil {
		log.Fatalf("Some error occurred in LAND. Err: %s", err)
	}
	land = l

	s, err := SALE.NewMain(saleAddress, client)
	if err != nil {
		log.Fatalf("Some error occurred in SALE. Err: %s", err)
	}
	sale = s

	// CHECK IF LOADED

	tokenName, err := token.Name(callOpts)
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}
	fmt.Printf("Token Loaded: %s\n", tokenName)

	landName, err := land.Name(callOpts)
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}
	fmt.Printf("Land Loaded: %s\n", landName)

	saleTokenName, err := sale.TOKEN(callOpts)
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}
	fmt.Printf("Sale Loaded with Token: %s\n", saleTokenName)

	// SUCCESS

	fmt.Println("🚀 Contracts are loaded!")
}
