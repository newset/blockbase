package ton

import (
	"crypto/ed25519"

	"github.com/tonkeeper/tongo/wallet"
)

func AddressFromSeed(seed string, version wallet.Version, dev bool) string {
	pk, _ := wallet.SeedToPrivateKey(seed)
	publicKey := pk.Public().(ed25519.PublicKey)
	address, _ := wallet.GenerateWalletAddress(publicKey, version, 0, nil)
	return address.ToHuman(true, dev)
}
