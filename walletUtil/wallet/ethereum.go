package wallet

import "math/big"

type Ethereum struct {
	AssetName string
}

func (e *Ethereum) GetLatestBlockNum() (*big.Int, error) {
	return nil, nil
}
func (e *Ethereum) GetBalance(address string) (*big.Int, error) {
	return nil, nil
}

func (e *Ethereum) GetUnsignTx(senderAddress, recieverAddress string, value big.Int) (string, error) {
	return "", nil
}

func (e *Ethereum) SignTx(privateKey string, unsignTx string) (string, error) {
	return "", nil
}

func (e *Ethereum) SendTx(signTx string) (string, error) {
	return "", nil
}
