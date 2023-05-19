package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d task %s is running\n", i, name)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go task("A")
	go task("B")
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("func anonymous %d, is running now\n", i)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(5 * time.Second)
}
