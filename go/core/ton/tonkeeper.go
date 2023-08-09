package ton

import (
	"context"
	"log"

	"github.com/tonkeeper/tongo"
	"github.com/tonkeeper/tongo/contract/jetton"
	"github.com/tonkeeper/tongo/liteapi"
	"github.com/tonkeeper/tongo/wallet"
)

var (
	CurrentSecret = ""
	Testnet       = true // true: mainnet, false: testnet
	WalletVersion = wallet.V4R2
	client        *liteapi.Client
	currentWallet wallet.Wallet
	Net           liteapi.Option
)

func CreateClient(isTest bool) {
	if isTest {
		_client, _ := liteapi.NewClientWithDefaultTestnet()
		client = _client
		return
	}

	_client, _ := liteapi.NewClientWithDefaultMainnet()
	client = _client
}

func GetTonByKeeper(mnemonic string, isTest bool) {
	_wallet := SetupAccount(mnemonic, client)
	currentWallet = _wallet
}

func GetJetton(token string) {

	master := tongo.MustParseAccountID(token)
	j := jetton.New(master, client)

	d, _ := j.GetDecimals(context.Background())
	log.Println("decimals: ", d)

	account := currentWallet.GetAddress()

	b, err := j.GetBalance(context.Background(), account)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("balance: ", b)

	jw, _ := j.GetJettonWallet(context.Background(), account)
	log.Println("jetton wallet: ", jw.String())
}

func TransferToken(to string, amount float64) {

}

// 验证jetton wallet和ton wallet的地址是否一致， 由于sdk的isTest参数定义不一致
func TestAccount() {
	// log.Println("test account: ", accountAddress)
	// account := tongo.MustParseAccountID(accountAddress)
	// log.Println("account: ", account)
	// b, err := client.GetBalance(context.Background(), account)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("balance: ", b)
}
