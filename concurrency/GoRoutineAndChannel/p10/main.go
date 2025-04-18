package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func(c chan int) {
		time.Sleep(5 * 1e9)
		c <- 10
		fmt.Println("sent", 10)
		// value recieved from channel
	}(c)

	fmt.Println("sending", 10)
	// c <- 10 // putting 10 on the channel
	fmt.Println("received", <-c)
}
