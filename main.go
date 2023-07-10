package main

import "fmt"

func plus(a, b int, potato string) (int, string) {
	return a + b, potato
}

func main() {
	fmt.Println("감자 코인을 만들자!")
	ans, name := plus(2, 2, "potato")
	fmt.Println(ans, name)
}
