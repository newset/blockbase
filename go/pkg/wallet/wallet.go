package wallet

import (
	"math/big"
)

// wallet struct
type Wallet struct {
	// 钱包ID
	Id int
	// 助记词
	Mnemonic string
	// 私钥
	PrivateKey string
	// 公钥
	PublicKey string
	// 地址
	Address string
	// 密码
	Password string
	// 余额
	Balance *big.Int
}

type Token struct {
	Id *int
	// 地址
	Address string

	decimals int

	Balance *big.Int
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
}
