package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go Publ(ch)
	go Reader(ch, &wg)
	wg.Wait()
}

func Reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Reveived %d\n", x)
		wg.Done()
	}
}

func Publ(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
