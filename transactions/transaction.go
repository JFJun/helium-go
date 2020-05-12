package transactions

type ITransaction interface {
	Serialize() ([]byte, error)
	BuildTransaction() ([]byte, error)
	SetSignature(sig []byte)
}
