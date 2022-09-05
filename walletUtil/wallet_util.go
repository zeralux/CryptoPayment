package walletUtil

import (
	"errors"
	"github.com/zeralux/CryptoPayment/walletUtil/model"
	"github.com/zeralux/CryptoPayment/walletUtil/wallet"
	"math/big"
)

// 手動加入
var accountWallets = map[string]wallet.AccountWallet{
	string(model.Ethereum): new(wallet.EthereumWallet),
}

// 手動加入
var uxtoWallets = map[string]wallet.UxtoWallet{
	model.BitCoin: new(wallet.BitCoinWallet),
}
var chainOperations = map[string]wallet.ChainOperation{}

func init() {
	for chainName, wallet := range accountWallets {
		chainOperations[chainName] = wallet
	}
	for chainName, wallet := range uxtoWallets {
		chainOperations[chainName] = wallet
	}
}

func GetLatestBlockNumber(chainName model.ChainName) (*big.Int, error) {
	if wallet, ok := chainOperations[string(chainName)]; ok {
		return wallet.GetLatestBlockNumber()
	}
	return nil, errors.New("wallet not found")
}

// 單筆轉帳
// 一次執行單一幣種
func Transfer(transferReq model.TransferReq) (string, error) {
	token := transferReq.Token
	if wallet, ok := accountWallets[string(token.ChainName)]; ok {
		return transferAccountValue(wallet, transferReq)
	}
	if wallet, ok := uxtoWallets[string(token.ChainName)]; ok {
		return transferUxtoValue(wallet, transferReq)
	}
	return "", errors.New("wallet not found")
}

// 單筆
func transferAccountValue(accountWallet wallet.AccountWallet, transferReq model.TransferReq) (string, error) {
	transferParam := model.TransferParam{
		FromAddress: transferReq.FromAddress,
		ToAddress:   transferReq.ToAddress,
		Value:       transferReq.Value,
	}
	unsignTx, _ := accountWallet.GetUnsignTx(transferParam)
	signTx, _ := accountWallet.SignTx(transferReq.PrivateKey, *unsignTx)
	return accountWallet.SendTx(*signTx)
}

// 單筆
func transferUxtoValue(uxtoWallet wallet.UxtoWallet, transferReq model.TransferReq) (string, error) {
	transferParams := []model.TransferParam{
		{
			FromAddress: transferReq.FromAddress,
			ToAddress:   transferReq.ToAddress,
			Value:       transferReq.Value,
		},
	}
	privateKeys := []string{transferReq.PrivateKey}
	unsignTx, _ := uxtoWallet.GetUnsignTx(transferParams)
	signTx, _ := uxtoWallet.SignTx(privateKeys, *unsignTx)
	return uxtoWallet.SendTx(*signTx)
}
