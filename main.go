package main

import (
	"fmt"
	"sync"
	"time"
)

func For(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch <- i
		wg.Done()
		time.Sleep(1 * time.Second)

	}
	close(ch)
}

func For2(ch chan int, wg1 *sync.WaitGroup) {
	for i := 10; i < 20; i++ {
		ch <- i
		wg1.Done()
		time.Sleep(1 * time.Second)

	}
	close(ch)
}

func main() {

	ch := make(chan int)
	ch2 := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(20)

	go For(ch, &wg)
	go For2(ch2, &wg)

	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}

	for z := range ch2 {
		fmt.Printf("Received %d\n", z)
	}

	wg.Wait()
}
