package model

import "math/big"

type Token struct {
	ChainName       ChainName
	AssetName       string
	ContractAddress string
	Decimal         big.Int
}
