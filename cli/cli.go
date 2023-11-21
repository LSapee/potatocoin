package cli

import (
	"flag"
	"fmt"
	"github.com/LSapee/potatocoin/explorer"
	"github.com/LSapee/potatocoin/rest"
	"os"
	"runtime"
)

func usage() {
	fmt.Printf("Welcome to 감자 코인\n\n")
	fmt.Printf("Please use the follwing flags:\n\n")
	fmt.Printf("-port=4000: 		Set the PORT of hte server\n")
	fmt.Printf("-mode=rest:		Choose between 'html' and 'rest'\n")
	runtime.Goexit()
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}
	//기본값이 필요함
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		//start rest api
		rest.Start(*port)
	case "html":
		//start html
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
