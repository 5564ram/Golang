package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(10 * 1e9)
		fmt.Println("received", <-c) // value recieved from channel
	}()
	fmt.Println("sending", 10)
	c <- 10 // putting 10 on the channel
	fmt.Println("sent", 10, "to channel len :", len(c))
}

// If "received", <-c is not printed, then:

// The channel is empty (len(c) == 0).
// Why? Because c <- 10 only completes if a receiver has already extracted the value.
// The problem is that the main goroutine exits before the receiver actually prints.
