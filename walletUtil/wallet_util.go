package walletUtil

import (
	"errors"
	"github.com/zeralux/gin/walletUtil/interfaces"
	"github.com/zeralux/gin/walletUtil/model"
	"github.com/zeralux/gin/walletUtil/util"
)

var wallets = []interfaces.Wallet{
	new(util.EthereumWallet),
	new(util.BitCoinWallet),
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
			if supportTokens[supportToken.GetId()] != nil {
				continue
			}
			supportTokens[supportToken.GetId()] = wallet
		}
	}
}

// 單筆轉帳
// 一次執行單一幣種
func Transfer(tokenId model.TokenId, transferReq model.TransferReq) (string, error) {
	wallet := supportTokens[tokenId]
	switch wallet.(type) {
	case interfaces.AccountWallet:
		accountWallet := wallet.(interfaces.AccountWallet)
		return transferAccountValue(accountWallet, &transferReq)
	case interfaces.UxtoWallet:
		return "", errors.New("util not found")
	default:
		return "", errors.New("util not found")
	}
}

// 單筆
func transferAccountValue(accountWallet interfaces.AccountWallet, transferReq *model.TransferReq) (string, error) {
	unsignTx, _ := accountWallet.GetUnsignTx(transferReq.SenderAddress, transferReq.RecieverAddress, transferReq.Value)
	signTx, _ := accountWallet.SignTx(transferReq.PrivateKey, unsignTx)
	return accountWallet.SendTx(signTx)
}
