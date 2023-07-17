package main

import (
	"fmt"
	"github.com/LSapee/potatocoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("감자 블록")
	chain.AddBlock("감자1 블록3")
	chain.AddBlock("감자2 블록2")
	chain.AddBlock("감자3 블록1")
	for _, block := range chain.AllBlocks() {
		fmt.Println("Data : ", block.Data)
		fmt.Println("Hash : ", block.Hash)
		fmt.Println("PrevHash : ", block.PrevHash)
	}

}
