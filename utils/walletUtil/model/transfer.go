package model

import "math/big"

type Transfer struct {
	fromAddress string
	toAddress   string
	value       big.Int
}
