package interfaces

import "github.com/zeralux/gin/utils/walletUtil/model"

type AccountWallet interface {
	Wallet
	accountTransfer
}

type accountTransfer interface {
	GetUnsignTx(transfer model.Transfer) (string, error)
	SignTx(privateKey string, unsignTx string) (string, error)
	SendTx(signTx string) (string, error)
}
