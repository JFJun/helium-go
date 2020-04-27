package crypto

import "testing"

func TestGenerateKey(t *testing.T) {
	ks := New(1)
	ks.GenerateKey()
}
