package core

import (
	"github.com/dabankio/wallet-core/bip39"
)

// NewEntropy will create random entropy bytes
func NewEntropy(bits int) (entropy []byte, err error) {
	return bip39.NewEntropy(bits)
}
