package model

import (
	"math/big"
)

type TransferReq struct {
	FromAddress string
	ToAddress   string
	Value       *big.Int
	PrivateKey  string
}