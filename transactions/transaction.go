package transactions

import (
	"encoding/hex"
	"fmt"
	"github.com/JFJun/helium-go/protos"
	"github.com/btcsuite/btcutil/base58"
	"github.com/golang/protobuf/proto"
)

//构建交易

func BuildPaymentV1Tx(from, to string, amount, fee, nonce uint64, sig []byte) ([]byte, error) {

	fromB := base58.Decode(from)
	payer := make([]byte, 33)
	copy(payer, fromB[1:len(fromB)-4])
	toB := base58.Decode(to)
	payee := make([]byte, 33)
	copy(payee, toB[1:len(toB)-4])

	//pay1,_:=base64.StdEncoding.DecodeString("ATEzTThkVWJ4eW1FM3h0aUFYc3pSa0dNbWV6TWhCUzhMaTd3RXNNb2pMZGI0U2R4YzR3Yw==")
	//pee1,_:=base64.StdEncoding.DecodeString("ATE0OGQ4S1RSY0tBNUpLUGVrQmNLRmQ0S2Z2cHJ2RlJwakd0aXZodG1SbW5aOE1GWW5QMw==")
	//fmt.Println(pay1)
	//fmt.Println(len(pay1))
	//fmt.Println(pee1)
	//fmt.Println(len(pee1))
	v1 := &protos.BlockchainTxnPaymentV1{
		Payer:     payer,
		Payee:     payee,
		Amount:    amount,
		Fee:       fee,
		Nonce:     nonce,
		Signature: nil,
	}

	//if sig != nil {
	//	v1.Signature = sig
	//}
	s, _ := hex.DecodeString("0a2101351a71c22fefec2231936ad2826b217ece39d9f77fc6c49639926299c38692951221019c659d723cc1e810a72e78f7deaf4736a87f10ef8fcfc80100b53327e7ee49a4180a2801")
	fmt.Println(s)
	fmt.Println(len(s))
	ss, _ := proto.Marshal(v1)
	fmt.Println(ss)
	fmt.Println(len(ss))

	txnPV1 := &protos.BlockchainTxn_Payment{Payment: v1}
	txn := &protos.BlockchainTxn{}
	txn.SetBlockTxn(txnPV1)

	return proto.Marshal(txn)
}
