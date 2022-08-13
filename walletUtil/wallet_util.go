package walletUtil

import (
	"errors"
	"github.com/zeralux/gin/walletUtil/interfaces"
	"github.com/zeralux/gin/walletUtil/model"
	"github.com/zeralux/gin/walletUtil/wallet"
)

var wallets = []interfaces.Wallet{
	new(wallet.Ethereum),
	new(wallet.BitCoin),
}

var supportTokens map[model.TokenId]interfaces.Wallet

// restful load static data
func init() {
	for _, wallet := range wallets {
		// check
		switch wallet.(type) {
		case interfaces.AccountWallet:
		case interfaces.UxtoWallet:
		default:
			continue
		}
		for _, supportToken := range wallet.GetSupportToken() {
			// check
			if supportTokens[supportToken.TokenId()] != nil {
				continue
			}
			supportTokens[supportToken.TokenId()] = wallet
		}
	}
}

// 單筆
func Transfer(transferData model.TransferData) (string, error) {
	tokenId := transferData.Token.TokenId()
	wallet := supportTokens[tokenId]
	switch wallet.(type) {
	case interfaces.AccountWallet:
		accountWallet := wallet.(interfaces.AccountWallet)
		return transferAccount(accountWallet, &transferData)
	case interfaces.UxtoWallet:
		return "", errors.New("wallet not found")
	default:
		return "", errors.New("wallet not found")
	}
}

// 單筆
func transferAccount(accountWallet interfaces.AccountWallet, transferData *model.TransferData) (string, error) {
	unsignTx, _ := accountWallet.GetUnsignTx(transferData.SenderAddress, transferData.RecieverAddress, transferData.Value)
	signTx, _ := accountWallet.SignTx(transferData.PrivateKey, unsignTx)
	return accountWallet.SendTx(signTx)
}
