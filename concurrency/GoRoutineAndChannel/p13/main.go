package main

import (
	"fmt"
	"time"
)

// semaphore pattern
func main() {
	c := make(chan int)
	go func(c chan int) {
		fmt.Println("sending", 10)
		time.Sleep(10 * 1e9)
		fmt.Println("send 10") // value recieved from channel
		c <- 10
	}(c)
	fmt.Println("received", <-c) // 10
}
