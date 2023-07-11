package trx

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-demo/core"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	trxProtoCore "github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/golang/protobuf/proto"
)

const addressPrefix = 0x41

const symbol = "TRX"

const TRX_MESSAGE_HEADER = "\x19TRON Signed Message:\n32"

// it should be: '\x15TRON Signed Message:\n32';
const ETH_MESSAGE_HEADER = "\x19Ethereum Signed Message:\n32"

// trx key derivation service
type trx struct{ core.CoinInfo }

// 签名hash
func SignHash(hash string, privateKeyHex string) (signature string) {
	txRawBytes, err := hex.DecodeString(hash)
	if err != nil {
		return "hex error " + err.Error()
	}
	tx := new(trxProtoCore.Transaction)
	err = proto.Unmarshal(txRawBytes, tx)
	if err != nil {
		return "transaction error" + err.Error()
	}

	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return "raw data error" + err.Error()
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKeyECDSA, _ := crypto.HexToECDSA(privateKeyHex)

	txHash := sha256.Sum256(rawData)
	sign, _ := crypto.Sign(txHash[:], privateKeyECDSA)
	tx.Signature = append(tx.Signature, sign)
	txSigBytes, _ := proto.Marshal(tx)
	return hex.EncodeToString(txSigBytes)
}

func SignString(message string, key string) (signature string) {

	fullMessage := fmt.Sprintf("\x19TRON Signed Message:\n%d%s", 32, message)
	hashData := crypto.Keccak256Hash([]byte(fullMessage))

	privateKeyECDSA, _ := crypto.HexToECDSA(key)
	sign, _ := crypto.Sign(hashData.Bytes(), privateKeyECDSA)

	// https://stackoverflow.com/a/69771013
	sign[64] += 27
	return hexutil.Encode(sign)
}

// 签名交易
func (c *trx) SignTransaction() {

}

func (c *trx) VerifyTransaction() {
	c.SignTransaction()
}
