package util

import (
	"github.com/zeralux/CryptoPayment/walletUtil/interfaces"
	"github.com/zeralux/CryptoPayment/walletUtil/model"
	"math/big"
)

type BitCoinWallet struct {
	token *model.Token
}

func (e *BitCoinWallet) NewWallet(token *model.Token) interfaces.UxtoWallet {
	return &BitCoinWallet{token: token}
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

func (e *BitCoinWallet) GetUnsignTx(transferParams []*model.TransferParam) (*model.UnsignTx, error) {
	return nil, nil
}

func (e *BitCoinWallet) SignTx(privateKey []string, unsignTx *model.UnsignTx) (*model.SignTx, error) {
	return nil, nil
}

func (e *BitCoinWallet) SendTx(signTx *model.SignTx) (string, error) {
	return "", nil
}
