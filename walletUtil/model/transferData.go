package model

import (
	"math/big"
)

type TransferData struct {
	Token           *Token
	SenderAddress   string
	RecieverAddress string
	Value           *big.Int
	PrivateKey      string
}
