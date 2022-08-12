package walletUtil

import (
	"math/big"
	"testing"
)

func TestNewWallet(t *testing.T) {
	walletUtil := NewWalletUtil("ETH")
	walletUtil.Transfer("aa", "bbb", big.NewInt(10))
}
