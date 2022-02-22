package main

import "math/big"

type Sale struct {
	Active bool
	Price  *big.Int
}

type Coordinates struct {
	x *big.Int
	y *big.Int
}