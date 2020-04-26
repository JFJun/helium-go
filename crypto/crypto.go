package crypto

import "crypto/ecdsa"

/*
Helium 支持ed25519和NIST p-256两种非对称加密格式
*/

func GenerateKey() {

	ecdsa.GenerateKey()
}
