package wallet

import (
	"github.com/zeralux/CryptoPayment/walletUtil/model"
	"math/big"
)

type EthereumWallet struct {
	token model.Token
}

func (e *EthereumWallet) NewWallet(token model.Token) AccountWallet {
	return &EthereumWallet{token: token}
}

func (e *EthereumWallet) GetChainName() model.ChainName {
	return model.Ethereum
}

func (e *EthereumWallet) GetLatestBlockNumber() (*big.Int, error) {
	return nil, nil
}

func (e *EthereumWallet) GetBalance(address string) (*big.Int, error) {
	return nil, nil
}

func (e *EthereumWallet) GetUnsignTx(transferParam model.TransferParam) (*model.UnsignTx, error) {
	return nil, nil
}

func (e *EthereumWallet) SignTx(privateKey string, unsignTx model.UnsignTx) (*model.SignTx, error) {
	return nil, nil
}

func (e *EthereumWallet) SendTx(signTx model.SignTx) (string, error) {
	return "", nil
}
