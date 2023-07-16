package wallet

import (
	"blockbase/core/trx"
)

//export SignTron
func SignTron(message string, privateKey string) string {
	sig := trx.SignString(message, privateKey)
	return sig
}
