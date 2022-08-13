package model

type Token struct {
	id              TokenId
	contractAddress string
}

type TokenId struct {
	chainName ChainName
	assetName AssetName
}

func (t *Token) GetId() TokenId {
	return t.id
}
