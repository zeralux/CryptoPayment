package interfaces

import (
	"math/big"
)

type Wallet interface {
	GetLatestBlockNum() (*big.Int, error)
	GetBalance(address string) (*big.Int, error)
}
