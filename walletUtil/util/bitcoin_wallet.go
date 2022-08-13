package util

import (
	"github.com/zeralux/gin/walletUtil/model"
	"math/big"
)

type BitCoinWallet struct {
}

func (e *BitCoinWallet) GetSupportToken() []*model.Token {
	return nil
}

func (e *BitCoinWallet) GetLatestBlockNum() (*big.Int, error) {
	return nil, nil
}

func (e *BitCoinWallet) GetBalance(address string) (*big.Int, error) {
	return nil, nil
}

func (e *BitCoinWallet) GetUnsignTx(senders []*model.Sender, recievers []*model.Reciever) (string, error) {
	return "", nil
}

func (e *BitCoinWallet) SignTx(privateKey []string, unsignTx string) (string, error) {
	return "", nil
}

func (e *BitCoinWallet) SendTx(signTx string) (string, error) {
	return "", nil
}
