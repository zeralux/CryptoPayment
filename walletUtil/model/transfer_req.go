package model

import (
	"math/big"
)

type TransferReq struct {
	Token       Token
	FromAddress string
	ToAddress   string
	Value       *big.Int
	PrivateKey  string
}
