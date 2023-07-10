package trx

import (
	"crypto/sha256"
	"encoding/hex"
	"blockbase/core"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	trxProtoCore "github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/golang/protobuf/proto"
)

const addressPrefix = 0x41

const symbol = "TRX"

// trx key derivation service
type trx struct{ core.CoinInfo }

// 签名hash
func SignHash(hash string, privateKeyHex string) (signature string) {
	txRawBytes, err := hex.DecodeString(hash)
	if err != nil {
		return "hex error" + err.Error()
	}
	tx := new(trxProtoCore.Transaction)
	err = proto.Unmarshal(txRawBytes, tx)
	if err != nil {
		return "transaction error"
	}

	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return "raw data error"
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKeyECDSA, _ := crypto.HexToECDSA(privateKeyHex)

	txHash := sha256.Sum256(rawData)
	sign, _ := crypto.Sign(txHash[:], privateKeyECDSA)
	tx.Signature = append(tx.Signature, sign)
	txSigBytes, _ := proto.Marshal(tx)
	return hex.EncodeToString(txSigBytes)
}

// 签名交易
func (c *trx) SignTransaction() {

}

func (c *trx) VerifyTransaction() {
	c.SignTransaction()
}
