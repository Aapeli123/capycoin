package wallet

import "crypto/rsa"

type Wallet struct {
	privKey *rsa.PrivateKey
	PubKey  *rsa.PublicKey
	Address []byte
}

func GenKeyPair() {

}

func LoadKeyPair() *Wallet {
	return nil
}

func SaveKeyPair() {

}
