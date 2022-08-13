package model

type Token struct {
	id              TokenId
	contractAddress string
}

type TokenId struct {
	chainName ChainName
	assetName AssetName
}

func (t *Token) TokenId() TokenId {
	return t.id
}
