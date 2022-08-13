package interfaces

import (
	"github.com/zeralux/CryptoPayment/walletUtil/model"
	"math/big"
)

type Wallet interface {
	GetSupportToken() []*model.Token
	GetLatestBlockNum() (*big.Int, error)
	GetBalance(address string) (*big.Int, error)
}
