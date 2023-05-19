package main

import (
	"fmt"
	"sync"
	"time"
)

func Number(description string, number int, wg *sync.WaitGroup, startChan chan bool) {
	for i := 1; i < 10; i++ {
		fmt.Printf("the %s is running at %d, on: %d\n", description, number, i)
		time.Sleep(2 * time.Second)

		startChan <- true
		wg.Done()
	}
}

func main() {
	waitgroup := sync.WaitGroup{}
	waitgroup.Add(25)

	startChan := make(chan bool)

	go Number("A", 1, &waitgroup, startChan)

	<-startChan
	go func() {
		for i := 9; i < 21; i++ {
			fmt.Printf("the func anonymous %d is running Now!\n", i)
			waitgroup.Done()
		}
	}()
	waitgroup.Wait()
}
