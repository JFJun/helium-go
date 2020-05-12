package transactions

import (
	"encoding/base64"
	"fmt"
	"github.com/JFJun/helium-go/keypair"
	"testing"
)

/*
	compare: github.com/helium/helium-js/package/transactions/src/__test__/PaymentV1.spec.ts
*/

var (
	alice    = "148d8KTRcKA5JKPekBcKFd4KfvprvFRpjGtivhtmRmnZ8MFYnP3"
	bob      = "13M8dUbxymE3xtiAXszRkGMmezMhBS8Li7wEsMojLdb4Sdxc4wc"
	from     = keypair.NewAddressable(bob)
	to       = keypair.NewAddressable(alice)
	v1       = NewPaymentV1Tx(from, to, 10, 0, 1, nil)
	toAmount = map[string]uint64{alice: 10}
	v2       = NewPaymentV2Tx(from, toAmount, 0, 1, nil)
	kp       = keypair.NewKeypairFromHex(1, "72eb1995e90e8b7c0054dcf594f4822572eb1995e90e8b7c0054dcf594f48225")
)

func TestPaymentV1Tx_Serialize(t *testing.T) {
	v1.Fee = 3
	v1.SetSignature([]byte("bob`s signature"))
	data, err := v1.Serialize()
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	/*
		=== RUN   TestPaymentV1Tx_Serialize
		Ql0KIQE1GnHCL+/sIjGTatKCayF+zjnZ93/GxJY5kmKZw4aSlRIhAZxlnXI8wegQpy54996vRzaofxDvj8/IAQC1Myfn7kmkGAogAygBMg9ib2JgcyBzaWduYXR1cmU=
		--- PASS: TestPaymentV1Tx_Serialize (0.00s)
		PASS
	*/
}

func TestPaymentV2Tx_Serialize(t *testing.T) {
	v2.Fee = 3
	v2.SetSignature([]byte("bob`s signature"))
	data, err := v2.Serialize()
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	/*
		=== RUN   TestPaymentV2Tx_Serialize
		wgFfCiEBNRpxwi/v7CIxk2rSgmshfs452fd/xsSWOZJimcOGkpUSJQohAZxlnXI8wegQpy54996vRzaofxDvj8/IAQC1Myfn7kmkEAoYAyABKg9ib2JgcyBzaWduYXR1cmU=
		--- PASS: TestPaymentV2Tx_Serialize (0.00s)
		PASS
	*/
}

func TestPaymentV1Tx_SignTransaction(t *testing.T) {

	v1Tx, err := v1.BuildTransaction()
	if err != nil {
		panic(err)
	}
	sig, err := kp.Sign(v1Tx)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(sig))
	/*
		=== RUN   TestPaymentV1Tx_SignTransaction
		yxLonpII3WdNiNg99WaRlS623HzkxIqPM7Vvjr62JtFnZrSP4zudIvz6vP/U9arXIlDbiyvO5nfiNM6tPmuzBw==
		--- PASS: TestPaymentV1Tx_SignTransaction (0.00s)
		PASS
	*/
}
func TestPaymentV2Tx_SignTransaction(t *testing.T) {

	v2Tx, err := v2.BuildTransaction()
	if err != nil {
		panic(err)
	}
	sig, err := kp.Sign(v2Tx)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(sig))
	/*
		=== RUN   TestPaymentV2Tx_SignTransaction
		rzwaOYb0KsRIdV1wv/9l9YP4qwGdsqZmoqXHygPz0VSazs6Nk1wvk2JDijSNESD5qGSNzaNfhKAoEGeIz5CgAA==
		--- PASS: TestPaymentV2Tx_SignTransaction (0.00s)
		PASS
	*/
}
