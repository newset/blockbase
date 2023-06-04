package tron

import (
	"crypto/ecdsa"
	"encoding/hex"
	"strings"

	wallet "blockbase/pkg/wallet"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keys"
	"github.com/tyler-smith/go-bip39"
)

func pkToAddress(pk *ecdsa.PrivateKey) string {
	add := address.PubkeyToAddress(pk.PublicKey)
	return common.EncodeCheck(add)
}

/**
 * 通过助记词生成账户
 */
func GenerateTronAccountFromMnemonic(mnemonic string) (*wallet.Account, error) {

	private, _ := keys.FromMnemonicSeedAndPassphrase(mnemonic, "", 0)
	privateKeyHex := hex.EncodeToString(private.Serialize())
	pk := private.ToECDSA()

	account := &wallet.Account{
		Address:    pkToAddress(pk),
		Mnemonic:   mnemonic,
		Privatekey: privateKeyHex,
	}

	return account, nil
}

/**
 * 通过私钥生成账户
 */
func GenerateTronAccountFromPrivateKey(privateKey string) (*wallet.Account, error) {
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

	account := &wallet.Account{
		Address:    pkToAddress(pk),
		Privatekey: privateKeyHex,
	}
	return account, nil
}

func CreateNewTronAccount() (*wallet.Account, error) {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return GenerateTronAccountFromMnemonic(mnemonic)
}
