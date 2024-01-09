package main

import (
	"github.com/LSapee/potatocoin/cli"
	"github.com/LSapee/potatocoin/db"
)

func main() {
	//wallet.Wallet()
	defer db.Close()
	cli.Start()
}
