package transactions

import (
	"github.com/JFJun/helium-go/protos"
	"github.com/golang/protobuf/proto"
)

//构建交易

func BuildPaymentV1Tx(from, to string, amount, fee, nonce uint64, sig []byte) ([]byte, error) {
	payer := []byte(from)
	payee := []byte(to)
	v1 := protos.PaymentV1{
		Payer:     payer,
		Payee:     payee,
		Amount:    amount,
		Fee:       fee,
		Nonce:     nonce,
		Signature: nil,
	}
	if sig != nil {
		v1.Signature = sig
	}

	return proto.Marshal(&v1)
}

func BuildPaymentV2Tx(from string, to_amount map[string]uint64, fee, nonce uint64, sig []byte) ([]byte, error) {
	payer := []byte(from)
	var payments []*protos.Payment
	for to, amount := range to_amount {
		payee := []byte(to)
		payment := &protos.Payment{
			Payee:  payee,
			Amount: amount,
		}
		payments = append(payments, payment)
	}
	v2 := protos.PaymentV2{
		Payer:     payer,
		Payments:  payments,
		Fee:       fee,
		Nonce:     nonce,
		Signature: nil,
	}
	if sig != nil {
		v2.Signature = sig
	}
	return proto.Marshal(&v2)
}
