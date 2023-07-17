package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// 모든 데이터는 블록에만
type block struct {
	data     string
	hash     string
	prevHash string
}

// 블록체인
type blockchain struct {
	blocks []*block
}

var b *blockchain

// 1번만 실행되게 해주기 위해
var once sync.Once

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.blocks = append(b.blocks, createBlock("Genesis Block"))
		})
	}
	return b
}