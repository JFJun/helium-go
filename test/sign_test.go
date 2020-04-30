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

}
