package test

import (
	"encoding/base64"
	"fmt"
	"github.com/JFJun/helium-go/keypair"
	"github.com/JFJun/helium-go/transactions"
	"testing"
)

/*
	compare: github.com/helium/helium-js/integration_tests/tests/create_and_submit_payment.spec.ts
*/

var (
	alice    = "148d8KTRcKA5JKPekBcKFd4KfvprvFRpjGtivhtmRmnZ8MFYnP3"
	bob      = "13M8dUbxymE3xtiAXszRkGMmezMhBS8Li7wEsMojLdb4Sdxc4wc"
	from     = keypair.NewAddressable(bob)
	to       = keypair.NewAddressable(alice)
	v1       = transactions.NewPaymentV1Tx(from, to, 10, 30000, 1, nil)
	toAmount = map[string]uint64{alice: 10}
	v2       = transactions.NewPaymentV2Tx(from, toAmount, 0, 1, nil)
	kp       = keypair.NewKeypairFromHex(1, "72eb1995e90e8b7c0054dcf594f4822572eb1995e90e8b7c0054dcf594f48225")
)

func Test_CreateAndSubmitPayment1(t *testing.T) {
	v1Tx, err := v1.BuildTransaction()
	if err != nil {
		panic(err)
	}
	sig, err := kp.Sign(v1Tx)
	if err != nil {
		panic(err)
	}
	v1.SetSignature(sig)
	ser, err := v1.Serialize()
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(ser))
	/*
		=== RUN   Test_CreateAndSubmitPayment1
		QowBCiEBNRpxwi/v7CIxk2rSgmshfs452fd/xsSWOZJimcOGkpUSIQGcZZ1yPMHoEKcuePfer0c2qH8Q74/PyAEAtTMn5+5JpBgKKAEyQMsS6J6SCN1nTYjYPfVmkZUuttx85MSKjzO1b46+tibRZ2a0j+M7nSL8+rz/1PWq1yJQ24srzuZ34jTOrT5rswc=
		--- PASS: Test_CreateAndSubmitPayment1 (0.00s)
		PASS
	*/
}
func Test_CreateAndSubmitPayment2(t *testing.T) {
	v2Tx, err := v2.BuildTransaction()
	if err != nil {
		panic(err)
	}
	sig, err := kp.Sign(v2Tx)
	if err != nil {
		panic(err)
	}
	v2.SetSignature(sig)
	ser2, err := v2.Serialize()
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(ser2))
	/*
		=== RUN   Test_CreateAndSubmitPayment2
		wgGOAQohATUaccIv7+wiMZNq0oJrIX7OOdn3f8bEljmSYpnDhpKVEiUKIQGcZZ1yPMHoEKcuePfer0c2qH8Q74/PyAEAtTMn5+5JpBAKIAEqQK88GjmG9CrESHVdcL//ZfWD+KsBnbKmZqKlx8oD89FUms7OjZNcL5NiQ4o0jREg+ahkjc2jX4SgKBBniM+QoAA=
		--- PASS: Test_CreateAndSubmitPayment2 (0.00s)
		PASS
	*/

}
