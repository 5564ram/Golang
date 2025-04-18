package main

import (
	"fmt"
)

func tel(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // if this is ommitted: fatal error: all goroutines are asleep - deadlock!
}

func main() {
	i := 0
	ch := make(chan int)

	go tel(ch)

	for i = range ch {
		fmt.Printf("the counter is at %d\n", i)
	}
}
