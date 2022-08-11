package walletUtil

import (
	"errors"
	"github.com/zeralux/gin/walletUtil/interfaces"
	"github.com/zeralux/gin/walletUtil/model"
	"github.com/zeralux/gin/walletUtil/wallet"
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
		return nil, errors.New("new WalletUtil fail, not implement " + assetName + " wallet")
	}
	return WalletUtil, nil
}

func (w *WalletUtil) Transfer(transferData model.TransferData) (string, error) {
	switch w.wallet.(type) {
	case interfaces.AccountWallet:
		wallet := w.wallet.(interfaces.AccountWallet)
		wallet.GetUnsignTx(transferData.SenderAddress, transferData.RecieverAddress, transferData.Value)

	case interfaces.UxtoWallet:
		wallet := w.wallet.(interfaces.UxtoWallet)
		wallet.GetUnsignTx(make([]model.Sender, 0), make([]model.Reciever, 0))
	}
	return "", errors.New("wallet not found")
}

func (w *WalletUtil) TransferBatch(transferData []model.TransferData) (string, error) {
	switch w.wallet.(type) {
	case interfaces.AccountWallet:
		//wallet := w.wallet.(interfaces.AccountWallet)
		//wallet.GetUnsignTx(transferData.SenderAddress, transferData.ToAddress, transferData.Value)

	case interfaces.UxtoWallet:
		//wallet := w.wallet.(interfaces.UxtoWallet)
		//wallet.GetUnsignTx(make([]model.Sender, 0), make([]model.Reciever, 0))
	}
	return "", errors.New("wallet not found")
}

func (w *WalletUtil) getAccountWallet() interfaces.AccountWallet {
	return new(wallet.Ethereum)
}

func (w *WalletUtil) getUxtoWallet() interfaces.UxtoWallet {
	return nil
}
