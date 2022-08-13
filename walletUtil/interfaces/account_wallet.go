package interfaces

import "github.com/zeralux/gin/walletUtil/model"

type AccountWallet interface {
	NewWallet(token *model.Token) AccountWallet
	Wallet
	accountTransfer
}

type accountTransfer interface {
	GetUnsignTx(transferParam *model.TransferParam) (*model.UnsignTx, error)
	SignTx(privateKey string, unsignTx *model.UnsignTx) (*model.SignTx, error)
	SendTx(signTx *model.SignTx) (string, error)
}
