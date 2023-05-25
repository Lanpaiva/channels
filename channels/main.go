package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(6 * time.Millisecond)

			ch <- true
		}
	}()

	resp := <-ch

	defer fmt.Println(resp)
}
