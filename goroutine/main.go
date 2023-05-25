package main

import (
	"fmt"
	"sync"
	"time"
)

func Number(description string, number int, wg *sync.WaitGroup) {
	for i := 1; i < 10; i++ {
		fmt.Printf("the %s is running at %d, on: %d\n", description, number, i)
		time.Sleep(2 * time.Second)

		wg.Done()
	}
}

func main() {
	waitgroup := sync.WaitGroup{}
	waitgroup.Add(25)

	go Number("A", 1, &waitgroup)

	go func() {
		for i := 9; i < 21; i++ {
			fmt.Printf("the func anonymous %d is running Now!\n", i)
			waitgroup.Done()
		}
	}()
	waitgroup.Wait()
}
