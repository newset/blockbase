package wallet

import (
	"crypto/ecdsa"
	"math/big"
)

type Account struct {
	privatekey *ecdsa.PrivateKey

	Mnemonic   string
	Privatekey string
	Password   string
	Name       string
	Address    string
	Balance    *big.Int
}

type Token struct {
	Id *int
	// 地址
	Address string

	Decimals int

	Balance *big.Int

	Account int
}

type NFT struct {
	Id int
	// 合约地址
	Contract string
	// token id
	TokenId int
	// url
	Url string

	Standard string

	Account int
}
