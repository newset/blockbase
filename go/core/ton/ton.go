package ton

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
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

var api = func() *ton.APIClient {
	client := liteclient.NewConnectionPool()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := client.AddConnectionsFromConfigUrl(ctx, "https://ton-blockchain.github.io/testnet-global.config.json")
	if err != nil {
		panic(err)
	}

	return ton.NewAPIClient(client)
}()

func CreateAcccoun(secret string, version int) {

	address, err := TonutilsFromSeed(secret, wallet.V4R2)
	fmt.Println("address:", address, err)
}

// NewTonAccount creates a new ton account
// secret: 12 words mnemonic
// dev: testnet or mainnet
func NewTonAccount(secret string, ver wallet.Version, dev bool) *TonAccount {
	// serverUrL := Ternary(dev, TestnetURL, MainnetURL)

	words := strings.Split(secret, " ")
	_wallet, err := wallet.FromSeed(api, words, ver)

	if err != nil {
		log.Fatalln("FromSeed err:", err.Error())
		return nil
	}

	log.Println("wallet address:", _wallet.Address())

	return &TonAccount{
		Wallet: _wallet,
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

func TransferToken(token string, receiver string) {
	ctx := api.Client().StickyContext(context.Background())
	w, _ := wallet.FromSeed(api, strings.Split(CurrentSecret, " "), wallet.V4R2)
	log.Println("test wallet:", w.Address().String())

	// initialize ton api lite connection wrapper
	master := jetton.NewJettonMasterClient(api, address.MustParseAddr(token))

	receiverAddress := address.MustParseAddr(receiver)

	tokenWallet, err := master.GetJettonWallet(ctx, w.Address())
	if err != nil {
		log.Fatal(err)
	}

	tokenBalance, err := tokenWallet.GetBalance(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("token balance:", tokenBalance.String())

	amt := tlb.MustFromTON("12")
	transferPayload, err := tokenWallet.BuildTransferPayload(receiverAddress, amt, tlb.MustFromTON("0.5"), nil)
	if err != nil {
		panic(err)
	}

	msg := wallet.SimpleMessage(tokenWallet.Address(), tlb.MustFromTON("0.5"), transferPayload)

	// w.Transfer()
	err = w.Send(context.Background(), msg, true)

	if err != nil {
		panic(err)
	}
	fmt.Println("waiting for confirmation")
	_, block, err := w.SendWaitTransaction(ctx, msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("transfer tokn finished", block)
}
