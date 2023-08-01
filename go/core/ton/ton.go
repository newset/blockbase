package ton

import (
	"blockbase/core"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/jetton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

var (
	testnet = "https://ton.org/testnet-global.config.json"
	// mainnet = "https://ton-blockchain.github.io/global.config.json"
	mainnet = "https://ton.org/global-config.json"
)

type TonAccount struct {
	Wallet     *wallet.Wallet
	Secret     string
	Blockchain *core.Blockchain
	api        *ton.APIClient
	ctx        context.Context
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

// NewTonAccount creates a new ton account
// secret: 12 words mnemonic
// dev: testnet or mainnet
func NewTonAccount(secret string, dev bool) *TonAccount {
	serverUrL := Ternary(dev, testnet, mainnet)
	api, _ := setupClient(serverUrL)
	words := strings.Split(secret, " ")
	w, err := wallet.FromSeed(api, words, wallet.V4R2)

	if err != nil {
		log.Fatalln("FromSeed err:", err.Error())
		return nil
	}

	log.Println("wallet address:", w.Address())

	return &TonAccount{
		Wallet: w,
		Secret: secret,
		ctx:    context.Background(),
		api:    api,
		Blockchain: &core.Blockchain{
			Rpc:     serverUrL,
			Testnet: dev,
		},
	}
}

func (account *TonAccount) GetBalance() string {
	api := account.api
	ctx := account.ctx
	w := account.Wallet
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

	fmt.Println("balance:", balance)

	return balance.String()
}

func (account *TonAccount) GetTokenBalance(token string) string {
	tokenContract := address.MustParseAddr(token)
	master := jetton.NewJettonMasterClient(account.api, tokenContract)

	tokenWallet, err := master.GetJettonWallet(account.ctx, account.Wallet.Address())
	if err != nil {
		log.Fatal(err)
	}

	tokenBalance, err := tokenWallet.GetBalance(account.ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("token balance:", tokenBalance.String())

	return tokenBalance.String()
}

func (account *TonAccount) TransferToken(token string, receiver string, amount float32, memo string) string {
	tokenContract := address.MustParseAddr(token)
	master := jetton.NewJettonMasterClient(account.api, tokenContract)

	tokenWallet, err := master.GetJettonWallet(account.ctx, account.Wallet.Address())
	if err != nil {
		log.Fatal(err)
	}

	tokenBalance, err := tokenWallet.GetBalance(account.ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("token balance:", tokenBalance.String())

	return ""
}
