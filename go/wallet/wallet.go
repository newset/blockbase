package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"strings"

	"blockbase/core"

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
func GenerateTronAccountFromMnemonic(mnemonic string) (*core.Account, error) {

	private, _ := keys.FromMnemonicSeedAndPassphrase(mnemonic, "", 0)
	privateKeyHex := hex.EncodeToString(private.Serialize())
	pk := private.ToECDSA()

	account := &core.Account{
		// privatekey: pk,
		Address:    pkToAddress(pk),
		Mnemonic:   mnemonic,
		Privatekey: privateKeyHex,
	}

	return account, nil
}

/**
 * 通过私钥生成账户
 */
func GenerateTronAccountFromPrivateKey(privateKey string) (*core.Account, error) {
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

	account := &core.Account{
		// privatekey: pk,
		Address:    pkToAddress(pk),
		Privatekey: privateKeyHex,
	}
	return account, nil
}

func CreateNewTronAccount() (*core.Account, error) {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	account, err := GenerateTronAccountFromMnemonic(mnemonic)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return account, nil
}

func Verify() {
	pk := "04811f1b4c96b2f26d0ec6cc74a51386c62b4633c28bbb20a1f2a0b64e9368ff"
	account, _ := GenerateTronAccountFromPrivateKey(pk)
	log.Println("Address:", account.Address)
}
