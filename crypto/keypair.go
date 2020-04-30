package crypto

import (
	"encoding/hex"
	"github.com/JFJun/helium-go/utils"
	"github.com/btcsuite/btcutil/base58"
)

type Keypair struct {
	curve Curves //选择非对称加密方式
	PrivateKey
	version int
}

func New(version int) *Keypair {
	c := NewCurve(version)
	ks := new(Keypair)
	ks.curve = c
	ks.version = version
	return ks
}

func NewKeypairFromWif(version int, wif string) *Keypair {
	kp := New(version)
	//todo
	//kp.PrivateKey =
	return kp
}
func NewKeypairFromHex(version int, privHex string) *Keypair {
	kp := New(version)
	data, err := hex.DecodeString(privHex)
	if err != nil {
		return nil
	}
	kp.PrivateKey = data
	return kp
}
func (kp *Keypair) GenerateKey() (PrivateKey, PublicKey) {
	return kp.curve.GenerateKey()
}

func (kp *Keypair) CreateAddress(publicKey PublicKey) string {
	var (
		payload  []byte
		vpayload []byte
	)
	v := kp.curve.GetVersion() //曲线版本号=》 1->ed25519 0-> NIST p256
	payload = append(v, publicKey[:]...)
	version := []byte{0} //主网版本号
	vpayload = append(version, payload...)
	//double sha256
	checksum := utils.DoubleSha256(vpayload)[:4]
	vpayload = append(vpayload, checksum...)

	return base58.Encode(vpayload)
}

func (kp *Keypair) Sign(message []byte) []byte {
	return kp.PrivateKey.Sign(kp.version, message)
}
