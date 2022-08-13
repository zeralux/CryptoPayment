package wallet

import (
	"github.com/zeralux/gin/walletUtil/model"
	"math/big"
)

type BitCoin struct {
}

func (e *BitCoin) GetSupportToken() []*model.Token {
	return nil
}

func (e *BitCoin) GetLatestBlockNum() (*big.Int, error) {
	return nil, nil
}

func (e *BitCoin) GetBalance(address string) (*big.Int, error) {
	return nil, nil
}

func (e *BitCoin) GetUnsignTx(senders []*model.Sender, recievers []*model.Reciever) (string, error) {
	return "", nil
}

func (e *BitCoin) SignTx(privateKey []string, unsignTx string) (string, error) {
	return "", nil
}

func (e *BitCoin) SendTx(signTx string) (string, error) {
	return "", nil
}
