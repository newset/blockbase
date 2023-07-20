package main

import "C"

import "blockbase/core/trx"

func main() {

}

//export signTron
func signTron(txId *C.char, privateKey *C.char) *C.char {
	var sig = trx.SignString(C.GoString(txId), C.GoString(privateKey))
	return C.CString(sig)
}

//export version
func version() {
	print("version: 1.0.0")
}

// https://cloud.tencent.com/developer/article/1786332
