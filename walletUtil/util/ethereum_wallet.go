package util

import (
	"github.com/zeralux/gin/walletUtil/model"
	"math/big"
)

type EthereumWallet struct {
}

func (e *EthereumWallet) GetSupportToken() []*model.Token {
	return nil
}

func (e *EthereumWallet) GetLatestBlockNum() (*big.Int, error) {
	return nil, nil
}

func (e *EthereumWallet) GetBalance(address string) (*big.Int, error) {
	return nil, nil
}

func (e *EthereumWallet) GetUnsignTx(senderAddress, recieverAddress string, value *big.Int) (string, error) {
	return "", nil
}

func (e *EthereumWallet) SignTx(privateKey string, unsignTx string) (string, error) {
	return "", nil
}

func (e *EthereumWallet) SendTx(signTx string) (string, error) {
	return "", nil
}
