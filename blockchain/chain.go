package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/LSapee/potatocoin/db"
	"github.com/LSapee/potatocoin/utils"
	"sync"
)

// 모든 데이터는 블록에만

// 블록체인
type blockchain struct {
	//blocks []*Block
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain

// 1번만 실행되게 해주기 위해
var once sync.Once

func (b *blockchain) restore(data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(b)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			fmt.Printf("NewestHash : %s\n Height:%d\n", b.NewestHash, b.Height)
			// search for checkpoint on the db
			checkpoint := db.Checkpoints()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				// restore b from bytes
				fmt.Printf("Restoring...\n")
				b.restore(checkpoint)
			}
		})
	}
	fmt.Printf("NewestHash : %s\nHeight:%d", b.NewestHash, b.Height)
	return b
}
