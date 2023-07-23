package main

import "C"

import "blockbase/core/trx"

func main() {

}

//export TRON_signTron
func TRON_signTron(txId *C.char, privateKey *C.char) *C.char {
	var sig = trx.SignString(C.GoString(txId), C.GoString(privateKey))
	return C.CString(sig)
}

//export TRON_getAccount
func TRON_getAccount(privateKey string) *C.char {
	return C.CString("hello")
}

//export TRON_transfer
func TRON_transfer(privateKey string, to string, amount string) {

}

//export TRON_transfer20
func TRON_transfer20(privateKey string, to string, contract string) {

}

//export TRON_transfer21
func TRON_transfer21(privateKey string, to string, contract string, id string) {

}

//export TRON_call
func TRON_call(privateKey string, to string, amount string) {

}

//export TRON_query
func TRON_query(privateKey string, to string, amount string) {

}

//export ETH_signETH() {
func ETH_signETH() {

}

//export ETH_getAccount
func ETH_getAccount(privateKey string) *C.char {
	return C.CString("hello")
}

//export ETH_transfer
func ETH_transfer(privateKey string, to string, amount string) {

}

//export ETH_transfer20
func ETH_transfer20(privateKey string, to string, contract string) {

}

//export ETH_transfer21
func ETH_transfer21(privateKey string, to string, contract string, id string) {

}

//export ETH_call
func ETH_call(privateKey string, to string, amount string) {

}

//export ETH_query
func ETH_query(privateKey string, to string, amount string) {

}

//export TON_signETH() {
func TON_signETH() {

}

//export TON_getAccount
func TON_getAccount(privateKey string) *C.char {
	return C.CString("hello")
}

//export TON_transfer
func TON_transfer(privateKey string, to string, amount string) {

}

//export TON_transfer20
func TON_transfer20(privateKey string, to string, contract string) {

}

//export TON_transfer21
func TON_transfer21(privateKey string, to string, contract string, id string) {

}

//export TON_call
func TON_call(privateKey string, to string, amount string) {

}

//export TON_query
func TON_query(privateKey string, to string, amount string) {

}

//export version
func version() {
	print("version: 1.0.0")
}

// https://cloud.tencent.com/developer/article/1786332
