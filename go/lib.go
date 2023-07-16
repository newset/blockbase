package main

// #include <stdio.h>
// #include <errno.h>
import "C"
import (
	"blockbase/core/trx"
)

//export SignTron
func SignTron(message *C.char, privateKey *C.char) *C.char {
	sig := trx.SignString(C.GoString(message), C.GoString(privateKey))
	return C.CString(sig)
}

func main() {

}
