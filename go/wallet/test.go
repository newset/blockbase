package wallet

import (
	"blockbase/core/trx"
	"log"
)

//export Hello
func Hello(name string) string {
	return "hello " + name
}

//export Test
func Test() {
	// 生成助记词
	// entropy, _ := bip39.NewEntropy(128)
	// mnemonic, _ := bip39.NewMnemonic(entropy)
	mnemonic := "bracket hero joke uphold detail omit because absent diesel hand butter quantum"
	account, err := GenerateTronAccountFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Mnemonic:", account.Mnemonic)
	log.Println("Address:", account.Address)
	log.Println("PrivateKey:", account.Privatekey)

	log.Println("---")

	Verify()

	sig := trx.SignString("helloworld", account.Privatekey)
	log.Println("Signature:", sig)
	log.Println("Signature:", "0xb0993408799d7f977923c727eae1f08fa210b575783252c3b3adf557ab84d2200a720ec2297c9b8f868ac196977ce4c8ebcefa5eed5b5ea1d524523b875c11f31b")
}
