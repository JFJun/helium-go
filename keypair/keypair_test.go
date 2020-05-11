package keypair

import (
	"fmt"
	"testing"
)

var kp = New(1)

func TestKeypair_GenerateKey(t *testing.T) {
	priv, pub := kp.GenerateKey()
	fmt.Println(priv)
	fmt.Println(pub)
	fmt.Println(len(priv))
	fmt.Println(len(pub))
}
