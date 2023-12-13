package wallet

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/LSapee/potatocoin/utils"
	"math/big"
)

// hashedMessage , privateKey,signature은 모두 다른 곳에서 불러와서 각자 다른 방법으로 복구를 함
const (
	hashedMessage string = "c33084feaa65adbbbebd0c9bf292a26ffc6dea97b170d88e501ab4865591aafd"
	privateKey    string = "30770201010420110fa1076b89ed41dd993d10ffb40a54dc00fd2720c2c770e6524713068349a4a00a06082a8648ce3d030107a14403420004315997a9a2116c76462a423cd33b351be2a9aa9e11e9aeb0215f903b9b0c433255ea5db56154a26a811183dd049e6c3a526a95e5f9e48f9282c0b690c82aecac"
	signature     string = "e08cf51a25c9608d8a83c31f45deecd5334f2e6bf08600c14bbf639937f53729a3ffe96327063d73472c52ae33c70393456261ff93dbd65ca6e06cb1c6523c47"
)

func Start() {

	////ecdsa GO포함된 패키지 (Ellipc Curve Digital Signature Algorithm)
	//privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//
	//keyAsByes, err := x509.MarshalECPrivateKey(privateKey)
	//fmt.Printf("%x\n\n\n", keyAsByes)
	//
	//utils.HandleErr(err)
	//// Hash값이 String이라 다시 []byte값으로 변환
	//hashAsBytes, err := hex.DecodeString(hashedMessage)
	//utils.HandleErr(err)
	//r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	//signature := append(r.Bytes(), s.Bytes()...)
	//fmt.Printf("%x\n", signature)
	//utils.HandleErr(err)
	////fmt.Printf("R:%d\nS:%d\n", r, s)
	//
	////검증
	////ok := ecdsa.Verify(&privateKey.PublicKey, hashAsBytes, r, s)
	////
	////fmt.Println(ok)

	//복구 과정
	// privateKey 복구
	//인코딩이 재대로 된건지 확인
	privateByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	private, err := x509.ParseECPrivateKey(privateByte)
	utils.HandleErr(err)
	fmt.Println(private)
	//서명 복구
	sigBytes, err := hex.DecodeString(signature)
	utils.HandleErr(err)
	//시작부터 중간까지
	rBytes := sigBytes[:len(sigBytes)/2]
	//중간부터 끝까지
	sBytes := sigBytes[len(sigBytes)/2:]
	//fmt.Printf("%d\n\n%d\n\n%d\n\n", sigBytes, rBytes, sBytes)
	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
	fmt.Println(bigR, bigS)

	hashBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	ok := ecdsa.Verify(&private.PublicKey, hashBytes, &bigR, &bigS)
	fmt.Println(ok)

}
