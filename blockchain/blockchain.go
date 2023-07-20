package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// 모든 데이터는 블록에만
type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

// 블록체인
type blockchain struct {
	blocks []*Block
}

var b *blockchain

// 1번만 실행되게 해주기 위해
var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
	newBlock.calculateHash()
	fmt.Println("data 길이는 : ", len(data))
	fmt.Println("data : ", data)
	return &newBlock
}
func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}
