package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprint(rw, "hello from home")
}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(port, nil))
}
