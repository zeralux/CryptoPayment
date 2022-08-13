package model

import (
	"math/big"
)

type TransferReq struct {
	SenderAddress   string
	RecieverAddress string
	Value           *big.Int
	PrivateKey      string
}
