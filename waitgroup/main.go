package main

import (
	"fmt"
	"sync"
	"time"
)

func Thread(stage string, wg *sync.WaitGroup, ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Thread %s times:%d\n", stage, i)
		time.Sleep(1 * time.Second)
		wg.Done()
		ch <- i
	}
	close(ch)
	wg.Wait()
}

func main() {
	waitgroup := sync.WaitGroup{}
	waitgroup.Add(10)
	ch := make(chan int)

	go Thread("A", &waitgroup, ch)

	for i := 0; 1 < 10; i++ {
		fmt.Printf("thrand anonymous running now on %d\n", i)
		time.Sleep(1 * time.Second)
		waitgroup.Done()

	}
	close(ch)

	waitgroup.Wait()
}
