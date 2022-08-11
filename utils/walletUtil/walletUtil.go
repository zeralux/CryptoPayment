package walletUtil

import (
	"errors"
	"github.com/zeralux/gin/utils/walletUtil/interfaces"
	"github.com/zeralux/gin/utils/walletUtil/wallet"
	"math/big"
)

type WalletUtil struct {
	assetName string
	wallet    interfaces.Wallet
}

func NewWalletUtil(assetName string) (*WalletUtil, error) {
	WalletUtil := &WalletUtil{assetName: assetName}
	if wallet := WalletUtil.getAccountWallet(); wallet != nil {
		WalletUtil.wallet = wallet
	}
	if wallet := WalletUtil.getUxtoWallet(); wallet != nil {
		WalletUtil.wallet = wallet
	}
	if WalletUtil.wallet == nil {
		return nil, errors.New("new WalletUtil fail, wallet not found")
	}
	return WalletUtil, nil
}

func (w *WalletUtil) SingleTransfer(senderAddress, recieverAddress string, value big.Int) (string, error) {
	switch w.wallet.(type) {
	case interfaces.AccountWallet:
		wallet := w.wallet.(interfaces.AccountWallet)
		wallet.GetUnsignTx(senderAddress, recieverAddress, value)

	case interfaces.UxtoWallet:
		wallet := w.wallet.(interfaces.UxtoWallet)
		wallet.GetUnsignTx()
	}
	return "", errors.New("wallet not found")
}

func (w *WalletUtil) getAccountWallet() interfaces.AccountWallet {
	return new(wallet.Ethereum)
}

func (w *WalletUtil) getUxtoWallet() interfaces.UxtoWallet {
	return nil
}
