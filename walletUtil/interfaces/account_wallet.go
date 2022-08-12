package interfaces

import (
	"math/big"
)

type AccountWallet interface {
	Wallet
	accountTransfer
}

type accountTransfer interface {
	GetUnsignTx(senderAddress, reeevierAddress string, value big.Int) (string, error)
	SignTx(privateKey string, unsignTx string) (string, error)
	SendTx(signTx string) (string, error)
}
