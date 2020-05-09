package test

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/JFJun/helium-go/crypto"
	"github.com/JFJun/helium-go/transactions"
	"github.com/tyler-smith/go-bip39"
	"testing"
)

func TestSign(t *testing.T) {
	kp := crypto.New(1)
	priv1, pub1 := kp.GenerateKey()
	from := kp.CreateAddress(pub1)
	fmt.Println(from)
	kp.PrivateKey = priv1
	_, pub2 := kp.GenerateKey()
	to := kp.CreateAddress(pub2)
	fmt.Println(to)
	a := "13M8dUbxymE3xtiAXszRkGMmezMhBS8Li7wEsMojLdb4Sdxc4wc"
	b := "148d8KTRcKA5JKPekBcKFd4KfvprvFRpjGtivhtmRmnZ8MFYnP3"
	paymentV1, err := transactions.BuildPaymentV1Tx(a, b, 10, 0, 1, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	sig := kp.PrivateKey.Sign(1, paymentV1)
	paymentV1Sig, err := transactions.BuildPaymentV1Tx(a, b, 10, 0, 1, sig)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(base64.StdEncoding.EncodeToString(paymentV1Sig))

}

func TestParsePrivate(t *testing.T) {

	entropy, err := bip39.EntropyFromMnemonic("indicate flee grace spirit trim safe access oppose void police calm energy")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(entropy)
	//entropy,err:=bip39.EntropyFromMnemonic("legal winner thank year wave sausage worth useful legal winner thank yellow")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(hex.EncodeToString(entropy))

}

func TestCreateAddress(t *testing.T) {

	kp2 := crypto.NewKeypairFromHex(1, "1f5b981baca0420259ab53996df7a8ce0e3549c6616854e7dff796304bafb6bf")
	priv2 := ed25519.NewKeyFromSeed(kp2.PrivateKey)

	pub2 := make(crypto.PublicKey, 32)

	copy(pub2, priv2[32:])
	fmt.Println(hex.EncodeToString(pub2))
	fmt.Println(pub2)
	fmt.Println(kp2.CreateAddress(pub2))
	pub := []byte{156, 101, 157, 114, 60, 193, 232, 16,
		167, 46, 120, 247, 222, 175, 71, 54,
		168, 127, 16, 239, 143, 207, 200, 1,
		0, 181, 51, 39, 231, 238, 73, 164}
	fmt.Println(hex.EncodeToString(pub))
	kp := crypto.New(1)
	address := kp.CreateAddress(pub)
	fmt.Println(address)

}

func Test_Sign(t *testing.T) {
	alice := "148d8KTRcKA5JKPekBcKFd4KfvprvFRpjGtivhtmRmnZ8MFYnP3"
	bob := "13M8dUbxymE3xtiAXszRkGMmezMhBS8Li7wEsMojLdb4Sdxc4wc"
	paymentV1, err := transactions.BuildPaymentV1Tx(bob, alice, 10, 0, 1, nil)
	fmt.Println("paymentV1:  ", paymentV1)
	if err != nil {
		fmt.Println(err)
		return
	}
	kp := crypto.New(1)
	seed, _ := hex.DecodeString("72eb1995e90e8b7c0054dcf594f4822572eb1995e90e8b7c0054dcf594f48225")
	kp.PrivateKey = seed
	sig := kp.PrivateKey.Sign(1, paymentV1)
	sign_paymentV1, _ := transactions.BuildPaymentV1Tx(bob, alice, 10, 1, 1, sig)
	fmt.Println(base64.StdEncoding.EncodeToString(sign_paymentV1))

}

func Test_Paymentv1(t *testing.T) {
	alice := "148d8KTRcKA5JKPekBcKFd4KfvprvFRpjGtivhtmRmnZ8MFYnP3"
	bob := "13M8dUbxymE3xtiAXszRkGMmezMhBS8Li7wEsMojLdb4Sdxc4wc"

	paymentV1, err := transactions.BuildPaymentV1Tx(bob, alice, 10, 0, 1, []byte("some signature"))
	if err != nil {
		panic(err)
	}
	fmt.Println("paymentV1:  ", paymentV1)

	fmt.Println(len(paymentV1))
	//50 14 115 111 109 101 32 115 105 103 110 97 116 117 114 101
	//bob:=[]byte("abcdefg")
	//fmt.Println(bob)
	//alice:=[]byte("123456")
	//fmt.Println(alice)
	//sig:=[]byte("test")
	//fmt.Println(sig)
	//bob:=[]byte{61,62,63,64,65,66,67}
	//alice:=[]byte{31,32,33,34,35,36}
	//sig:=[]byte{74,65,73,74}
	//v1:=&protos.BlockchainTxnPaymentV1{
	//	Payer:     bob,
	//	Payee:     alice,
	//	Amount:    10,
	//	Fee:       1,
	//	Nonce:     2,
	//	Signature: sig,
	//}
	//payment,err:=proto.Marshal(v1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(payment)
	//fmt.Println(len(payment))
}
