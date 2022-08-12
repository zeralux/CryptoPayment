package model

import "math/big"

type TransferData struct {
	SenderAddress   string
	RecieverAddress string
	Value           big.Int
}
