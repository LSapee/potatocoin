package main

import (
	"github.com/LSapee/potatocoin/cli"
	"github.com/LSapee/potatocoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
