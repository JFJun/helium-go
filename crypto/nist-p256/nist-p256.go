package nist_p256

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type NISTP256Curve struct {
	Version []byte
}

func (nc *NISTP256Curve) GenerateKey() ([]byte, []byte) {

	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	private := priv.D.Bytes()
	x := priv.PublicKey.X.Bytes()
	y := priv.PublicKey.Y.Bytes()

	pub := append(x, y...)
	return private, pub
}

func (nc *NISTP256Curve) GetVersion() []byte {
	return nc.Version
}

func (nc *NISTP256Curve) NewKeyFromSeed(seed []byte) ([]byte, []byte) {
	return nil, nil
}
