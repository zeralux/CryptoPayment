package wallet

import (
	"github.com/zeralux/CryptoPayment/walletUtil/model"
	"math/big"
)

type BitCoinWallet struct {
	token model.Token
}

func (e *BitCoinWallet) NewWallet(token model.Token) UxtoWallet {
	return &BitCoinWallet{token: token}
}

func (e *BitCoinWallet) GetChainName() model.ChainName {
	return model.BitCoin
}

func (e *BitCoinWallet) GetLatestBlockNumber() (*big.Int, error) {
	return nil, nil
}

func (e *BitCoinWallet) GetBalance(address string) (*big.Int, error) {
	return nil, nil
}

func (e *BitCoinWallet) GetUnsignTx(transferParams []model.TransferParam) (*model.UnsignTx, error) {
	return nil, nil
}

func (e *BitCoinWallet) SignTx(privateKey []string, unsignTx model.UnsignTx) (*model.SignTx, error) {
	return nil, nil
}

func (e *BitCoinWallet) SendTx(signTx model.SignTx) (string, error) {
	return "", nil
}
