package main

import (
	"github.com/LSapee/potatocoin/blockchain"
	"github.com/LSapee/potatocoin/cli"
	"github.com/LSapee/potatocoin/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()
}
