package interfaces

import (
	"github.com/zeralux/gin/utils/walletUtil/model"
)

type UxtoWallet interface {
	Wallet
	utxoTransfer
}

type utxoTransfer interface {
	GetUnsignTx(senders []model.Sender, recievers []model.Reciever) (string, error)
	SignTx(privateKey []string, unsignTx string) (string, error)
	SendTx(signTx string) (string, error)
}
