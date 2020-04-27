package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"golang.org/x/crypto/ed25519"
)

/*
Helium 支持ed25519和NIST p-256两种非对称加密格式
*/

type Curves interface {
	GenerateKey() (PrivateKey, PublicKey)
	GetVersion() []byte
}

type PrivateKey []byte
type PublicKey []byte

func NewCurve(version int) Curves {
	var c Curves
	if version == 0 {
		nc := &NISTP256Curve{version: []byte{byte(version)}}
		c = nc
	} else if version == 1 {
		ec := &Ed25519Curve{version: []byte{byte(version)}}
		c = ec
	}
	return c
}

type Ed25519Curve struct {
	version []byte
}

func (ec *Ed25519Curve) GenerateKey() (PrivateKey, PublicKey) {

	pub, priv, err := ed25519.GenerateKey(rand.Reader)

	if err != nil {
		panic(err)
	}

	return PrivateKey(priv), PublicKey(pub)
}
func (ec *Ed25519Curve) GetVersion() []byte {
	return ec.version
}

//=====================================================================================================================//
type NISTP256Curve struct {
	version []byte
}

func (nc *NISTP256Curve) GenerateKey() (PrivateKey, PublicKey) {

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
	return nc.version
}

func (priv PrivateKey) Sign(version int, message []byte) []byte {
	var data []byte
	if version == 1 {
		p := ed25519.PrivateKey(priv)
		data = ed25519.Sign(p, message)
	} else {
		//todo
	}
	return data
}
