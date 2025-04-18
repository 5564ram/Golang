package main

import (
	"fmt"
	"time"
)

// The pump function returns a channel that pumps out the sequence of numbers.
// producer - consumer pattern
func main() {
	suck(pump())
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 100000; i++ {
			ch <- i
		}
	}(ch)
	return ch
}

func suck(ch chan int) {
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
}
