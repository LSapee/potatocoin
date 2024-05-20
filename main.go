package main

import (
	"fmt"
	"time"
)

// chan<- or <-chan이 아니라 그냥 c chan int 하면 보내기,받기 둘다 가능함

// 채널로 값을 보내기 전용
func countToTen(c chan<- int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)
}

// 채널로 값을 받기 정용
func receive(c <-chan int) {
	for {
		a, ok := <-c
		if !ok {
			break
		}
		fmt.Println(a)
	}
}

func main() {
	// 채널 만들기
	c := make(chan int)
	go countToTen(c)
	receive(c)
	//wallet.Wallet()
	//defer db.Close()
	//cli.Start()
}
