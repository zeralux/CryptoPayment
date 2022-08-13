package walletUtil

import (
	"errors"
	"github.com/zeralux/CryptoPayment/walletUtil/interfaces"
	"github.com/zeralux/CryptoPayment/walletUtil/model"
	"github.com/zeralux/CryptoPayment/walletUtil/util"
)

var accountWallets = []interfaces.AccountWallet{
	new(util.EthereumWallet),
}
var uxtoWallets = []interfaces.UxtoWallet{
	new(util.BitCoinWallet),
}
var accountTokens = map[model.TokenId]interfaces.AccountWallet{}
var uxtoTokens = map[model.TokenId]interfaces.UxtoWallet{}

// restful load static data
func init() {
	for _, accountWallet := range accountWallets {
		for _, supportToken := range accountWallet.GetSupportToken() {
			accountTokens[supportToken.GetId()] = accountWallet.NewWallet(supportToken)
		}
	}
	for _, uxtoWallet := range uxtoWallets {
		for _, supportToken := range uxtoWallet.GetSupportToken() {
			uxtoTokens[supportToken.GetId()] = uxtoWallet.NewWallet(supportToken)
		}
	}
}

// 單筆轉帳
// 一次執行單一幣種
func Transfer(tokenId model.TokenId, transferReq model.TransferReq) (string, error) {
	if wallet, ok := accountTokens[tokenId]; ok {
		return transferAccountValue(wallet, &transferReq)
	}
	if wallet, ok := uxtoTokens[tokenId]; ok {
		return transferUxtoValue(wallet, &transferReq)
	}
	return "", errors.New("wallet not found")
}

// 單筆
func transferAccountValue(accountWallet interfaces.AccountWallet, transferReq *model.TransferReq) (string, error) {
	transferParam := &model.TransferParam{
		FromAddress: transferReq.FromAddress,
		ToAddress:   transferReq.ToAddress,
		Value:       transferReq.Value,
	}
	unsignTx, _ := accountWallet.GetUnsignTx(transferParam)
	signTx, _ := accountWallet.SignTx(transferReq.PrivateKey, unsignTx)
	return accountWallet.SendTx(signTx)
}

// 單筆
func transferUxtoValue(uxtoWallet interfaces.UxtoWallet, transferReq *model.TransferReq) (string, error) {
	transferParams := []*model.TransferParam{
		{
			FromAddress: transferReq.FromAddress,
			ToAddress:   transferReq.ToAddress,
			Value:       transferReq.Value,
		},
	}
	privateKeys := []string{transferReq.PrivateKey}
	unsignTx, _ := uxtoWallet.GetUnsignTx(transferParams)
	signTx, _ := uxtoWallet.SignTx(privateKeys, unsignTx)
	return uxtoWallet.SendTx(signTx)
}
