package main

import (
	"blockbase/core/ton"
	"blockbase/core/trx"
	"fmt"
)

var (
	// tronAdress      = "grpc.nile.trongrid.io:50051"
	// accountAddress  = "TPpw7soPWEDQWXPCGUMagYPryaWrYR5b3b"
	// EQCWaXAdYlJcPayZdeXOuQZ24u0-FJnSTaWMj8dTXrD5tTcc
	testnetMnemonic = ""
	// EQC5K5Xzk6D_scb4PFQ8iFe9r-Nak-Df-6iDn5Ia5-xR09BQ
	mainnetMnemonic = ""
	testTonToken    = "EQAkwg92CNna9Q_jRFTDIitdWgqp-dpIXX9j-MHj1r0W-tYL"
	mainTonToken    = "EQBynBO23ywHy_CgarY9NK9FTz0yDsG82PtcbSTQgGoXwiuA"
	// JUST EQBynBO23ywHy_CgarY9NK9FTz0yDsG82PtcbSTQgGoXwiuA
)

func getSign() {
	var txId = ""
	privateKey := "0x0000"
	sig, _ := trx.SignTransactionId(txId, privateKey)
	println(sig)
}

func init() {
	getSign()

	fmt.Println("init")
	fmt.Println("n", mainnetMnemonic)
	testAccount := ton.NewTonAccount(testnetMnemonic, true)
	mainAccount := ton.NewTonAccount(mainnetMnemonic, false)
	fmt.Println("account", testAccount)
	testAccount.GetBalance()
	mainAccount.GetBalance()
	// testAccount.GetTokenBalance(testTonToken)
	mainAccount.GetTokenBalance(mainTonToken)

	// useAccount(1, "0x0000", "TPpw7soPWEDQWXPCGUMagYPryaWrYR5b3b")
	// fmt.Println("new account", account)
}
