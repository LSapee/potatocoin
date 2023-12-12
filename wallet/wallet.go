package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/LSapee/potatocoin/utils"
)

func Start() {
	//ecdsa GO포함된 패키지 (Ellipc Curve Digital Signature Algorithm)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	//1. 메시지 생성
	message := "I love you"
	hashedMessage := utils.Hash(message)
	// Hash값이 String이라 다시 []byte값으로 변환
	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleErr(err)
	fmt.Printf("R:%d\nS:%d", r, s)
}
