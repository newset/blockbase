package ton

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/jetton"
	"github.com/xssnick/tonutils-go/ton/nft"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

var (
	// mainnet = "https://ton-blockchain.github.io/global.config.json"
	TestnetURL = "https://ton.org/testnet-global.config.json"
	MainnetURL = "https://ton.org/global-config.json"
	CurrentNet = Testnet
)

type TonAccount struct {
	Wallet  *wallet.Wallet
	Secret  string
	Address string
	api     *ton.APIClient
}

func Ternary[T any](condition bool, If, Else T) T {
	if condition {
		return If
	}
	return Else
}

func setupClient(rpc string) (*ton.APIClient, context.Context) {
	client := liteclient.NewConnectionPool()
	err := client.AddConnectionsFromConfigUrl(context.Background(), rpc)
	if err != nil {
		panic(err)
	}

	ctx := client.StickyContext(context.Background())

	// initialize ton api lite connection wrapper
	api := ton.NewAPIClient(client)

	return api, ctx
}

func CreateAcccoun(secret string, version int) {

	address, err := TonutilsFromSeed(secret, wallet.V4R2)
	fmt.Println("address:", address, err)
}

// NewTonAccount creates a new ton account
// secret: 12 words mnemonic
// dev: testnet or mainnet
func NewTonAccount(secret string, ver wallet.Version, dev bool) *TonAccount {
	serverUrL := Ternary(dev, TestnetURL, MainnetURL)
	api, _ := setupClient(serverUrL)
	words := strings.Split(secret, " ")
	w, err := wallet.FromSeed(api, words, ver)

	if err != nil {
		log.Fatalln("FromSeed err:", err.Error())
		return nil
	}

	log.Println("wallet address:", w.Address())

	return &TonAccount{
		Wallet: w,
		Secret: secret,
		api:    api,
	}
}

func (account *TonAccount) GetBalance() string {
	api := account.api
	w := account.Wallet
	ctx := context.Background()

	block, err := api.CurrentMasterchainInfo(ctx)
	if err != nil {
		log.Fatalln("CurrentMasterchainInfo err:", err.Error())
		return ""
	}
	balance, err := w.GetBalance(ctx, block)

	if err != nil {
		log.Fatalln("GetBalance err:", err.Error())
		return ""
	}

	return balance.String()
}

func (account *TonAccount) GetTokenBalance(token string) string {
	tokenContract := address.MustParseAddr(token)
	master := jetton.NewJettonMasterClient(account.api, tokenContract)
	ctx := context.Background()

	tokenWallet, err := master.GetJettonWallet(ctx, account.Wallet.Address())
	if err != nil {
		log.Fatal(err)
	}

	data, _ := master.GetJettonData(ctx)
	content := data.Content.(*nft.ContentSemichain)
	log.Println("data:", content)
	log.Println("	symbol:", content.GetAttribute("symbol"))

	if content.GetAttribute("decimals") != "" {
		log.Println("	decimals:", content.GetAttribute("decimals"))
	} else {
		log.Println("	decimals:", 9)
	}

	tokenBalance, err := tokenWallet.GetBalance(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("token balance:", tokenBalance.String())

	return tokenBalance.String()
}

func (account *TonAccount) TransferToken(token string, receiver string, amount float32, memo string) {
	tokenContract := address.MustParseAddr(token)
	master := jetton.NewJettonMasterClient(account.api, tokenContract)
	ctx := context.Background()

	tokenWallet, err := master.GetJettonWallet(ctx, account.Wallet.Address())
	if err != nil {
		log.Fatal(err)
	}

	tokenBalance, err := tokenWallet.GetBalance(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("token balance:", tokenBalance.String())

}
