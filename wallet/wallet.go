package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"github.com/LSapee/potatocoin/utils"
	"os"
)

const (
	fileName string = "potatoCoin.wallet"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func createPrivateKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)
}

func Wallet() *wallet {

	if w == nil {
		w = &wallet{}
		// wallet이 이미 존재하는가?
		if hasWalletFile() {
			// 존재 한다면 ->

		} else {
			// 없다면 ->
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}

	}
	return w
}
