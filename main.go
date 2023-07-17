package main

import (
	"crypto/sha256"
	"fmt"
)

// 모든 데이터는 블록에만
type block struct {
	data     string
	hash     string
	prevHash string
}

// 블록체인
type blockchain struct {
	blocks []block
}

/*
*
블록체인하는 기법?
다른 블록을 만들때 기존의 블록을 이용해서 만든다?
B1

	b1Hash = (data +"")

B2

	b2Hash = (data +b1Hash)

B3

	b3Hash = (data +b2Hash)

B1 의 b1Hash를 b1Hash1로 변경하면
B1

	b1Hash1 = (data+"")

B2

	b2Hash = (data+b1Hash) 였던 것에서 b1Hash라는게 없어져서 (data + b1Hash1)로 바뀌어야 한다.

B3

	b3Hash = (data+b2Hast) b2Hash는 바뀌지 않았기에 기존이랑 똑같다.
*/
func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		// blocks의 길이가 0보다 크면 이전 블록의 해쉬값을 가져온다.
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

// 블록 추가
func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, newBlock)
}

// 블록 리스트
func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data : %s\n ", block.data)
		fmt.Printf("Hash : %s\n ", block.hash)
		fmt.Printf("prevHash : %s\n ", block.prevHash)
	}
}

func main() {
	chain := blockchain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("감자")
	chain.addBlock("potato")
	chain.listBlocks()
	//for _, aByte := range "Genesis Block" {
	//	fmt.Printf("%b\n", aByte)
	//}
	//genesisBlock := block{"Genesis Block", "", ""}
	// sha256.Sum256은 []byte slice 형식을 받는다.
	//string을 []byte로 만들고 해싱을 한다
	//hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	// [size]byte를 다시 string으로 만들어서 저장한다.
	//HexHash := fmt.Sprintf("%x", hash)
	// 해당 string을 hash에 저장해주기
	//genesisBlock.hash = HexHash
	//fmt.Println(genesisBlock)
}
