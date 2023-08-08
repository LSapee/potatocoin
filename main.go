package main

import (
	"github.com/LSapee/potatocoin/explorer"
	"github.com/LSapee/potatocoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(5000)
}
