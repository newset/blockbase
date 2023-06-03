package tron

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keys"
	"github.com/tyler-smith/go-bip39"
)

type TronClient struct {
	Id   int
	Name string
	Url  string
}


type TronAccount struct {
	Address    string
	PrivateKey string
}

type Account struct {
	privatekey *ecdsa.PrivateKey

	Mnemonic   string
	Privatekey string
	Phrase     string
	Name       string
	Address    string
	Balance    *big.Int
}

func pkToAddress(pk *ecdsa.PrivateKey) string {
	add := address.PubkeyToAddress(pk.PublicKey)
	return common.EncodeCheck(add)
}

/**
 * 通过助记词生成账户
 */
func GenerateTronAccountFromMnemonic(mnemonic string) (*Account, error) {

	private, _ := keys.FromMnemonicSeedAndPassphrase(mnemonic, "", 0)
	privateKeyHex := hex.EncodeToString(private.Serialize())
	pk := private.ToECDSA()

	account := &Account{
		privatekey: pk,
		Address:    pkToAddress(pk),
		Mnemonic:   mnemonic,
		Privatekey: privateKeyHex,
	}

	return account, nil
}

/**
 * 通过私钥生成账户
 */
func GenerateTronAccountFromPrivateKey(privateKey string) (*Account, error) {
	privateKey = strings.TrimPrefix(privateKey, "0x")

	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	if len(privateKeyBytes) != common.Secp256k1PrivateKeyBytesLength {
		return nil, common.ErrBadKeyLength
	}

	private, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	pk := private.ToECDSA()
	privateKeyHex := hex.EncodeToString(private.Serialize())

	account := &Account{
		privatekey: pk,
		Address:    pkToAddress(pk),
		Privatekey: privateKeyHex,
	}
	return account, nil
}
