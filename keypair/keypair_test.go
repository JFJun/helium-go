package keypair

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var kp = New(Ed25519Version)

func TestKeypair_GenerateKey(t *testing.T) {
	priv, pub := kp.GenerateKey()
	fmt.Println(priv)
	fmt.Println(pub)
	fmt.Println(len(priv))
	fmt.Println(len(pub))
	fmt.Println(kp.CreateAddress(pub))

	kp2 := NewKeypairFromHex(Ed25519Version, hex.EncodeToString(priv))
	a := kp2.CreateAddressable()
	fmt.Println(a.GetAddress())

}

func TestKeypair_GenerateWifAndAddress(t *testing.T) {
	wif, address := kp.GenerateWifAndAddress()
	fmt.Println(wif, address)
	kp2 := NewKeypairFromWIF(Ed25519Version, wif)
	a := kp2.CreateAddressable()
	fmt.Println(wif, a.GetAddress())
}

func TestKeypair_Sign(t *testing.T) {
	priv, _ := kp.GenerateKey()
	kp.SetPrivateKey(priv)
	data, err := kp.Sign([]byte("test ed25519 signature"))
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
