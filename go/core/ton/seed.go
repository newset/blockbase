package ton

import (
	"blockbase/core"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha512"
	"errors"
	"fmt"
	"strings"

	"github.com/xssnick/tonutils-go/ton/wallet"
	"golang.org/x/crypto/pbkdf2"
)

const (
	_Iterations   = 100000
	_Salt         = "TON default seed"
	_BasicSalt    = "TON seed version"
	_PasswordSalt = "TON fast seed version"
)

const DefaultSubwallet = 698983191

func FromSeed(seed string, version wallet.Version) (*core.Wallet, error) {
	words := strings.Split(seed, " ")
	// validate seed
	if len(words) < 12 {
		return nil, fmt.Errorf("seed should have at least 12 words")
	}

	mac := hmac.New(sha512.New, []byte(seed))
	hash := mac.Sum(nil)

	p := pbkdf2.Key(hash, []byte(_BasicSalt), _Iterations/256, 1, sha512.New)
	if p[0] != 0 {
		return nil, errors.New("invalid seed")
	}

	k := pbkdf2.Key(hash, []byte(_Salt), _Iterations, 32, sha512.New)
	// k.Public().(ed25519.PublicKey)
	key := ed25519.NewKeyFromSeed(k)
	add, err := wallet.AddressFromPubKey(key.Public().(ed25519.PublicKey), version, DefaultSubwallet)
	fmt.Println("add from seed: ", add, err)
	return &core.Wallet{
		Address:  add.String(),
		Mnemonic: seed,
	}, nil
}
