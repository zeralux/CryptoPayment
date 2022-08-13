package model

import "math/big"

type TransferParam struct {
	FromAddress string
	ToAddress   string
	Value       *big.Int
}
