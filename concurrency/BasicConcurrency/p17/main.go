package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel1 <- "I'll print every 100ms"
			time.Sleep(time.Millisecond * 100)

		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel2 <- "I'll print every 1s"
			time.Sleep(time.Second * 1)

		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)
		default:
			fmt.Println("No channel is ready")
		}
	}
}

/*
The default case executes when every other send/receive operation is blocked.
But what happens if we have more than one send/receive operation ready at the
same time? Well, the select statement chooses one of them at random.
*/
