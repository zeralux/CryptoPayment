package wallet

import (
	"github.com/zeralux/CryptoPayment/walletUtil/model"
	"math/big"
)

// Operation
// 一般操作
type ChainOperation interface {
	GetChainName() model.ChainName
	GetLatestBlockNumber() (*big.Int, error)
	GetBalance(address string) (*big.Int, error)
}

// AccountWallet
// account模型錢包
type AccountWallet interface {
	NewWallet(token model.Token) AccountWallet
	ChainOperation
	accountOperation
}

// account模型轉帳
type accountOperation interface {
	GetUnsignTx(transferParam model.TransferParam) (*model.UnsignTx, error)
	SignTx(privateKey string, unsignTx model.UnsignTx) (*model.SignTx, error)
	SendTx(signTx model.SignTx) (string, error)
}

// UxtoWallet
// utxo模型錢包
type UxtoWallet interface {
	NewWallet(token model.Token) UxtoWallet
	ChainOperation
	utxoOperation
}

// utxo模型轉帳
type utxoOperation interface {
	GetUnsignTx(transferParams []model.TransferParam) (*model.UnsignTx, error)
	SignTx(privateKey []string, unsignTx model.UnsignTx) (*model.SignTx, error)
	SendTx(signTx model.SignTx) (string, error)
}
