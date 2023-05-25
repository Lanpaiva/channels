package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go publ(ch)
	reader(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Reveived %d\n", x)
	}
}

func publ(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
