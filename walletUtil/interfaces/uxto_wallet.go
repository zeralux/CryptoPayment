package interfaces

import (
	"github.com/zeralux/gin/walletUtil/model"
)

type UxtoWallet interface {
	NewWallet(token *model.Token) UxtoWallet
	Wallet
	utxoTransfer
}

type utxoTransfer interface {
	GetUnsignTx(transferParams []*model.TransferParam) (*model.UnsignTx, error)
	SignTx(privateKey []string, unsignTx *model.UnsignTx) (*model.SignTx, error)
	SendTx(signTx *model.SignTx) (string, error)
}
