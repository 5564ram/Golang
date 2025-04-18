package main

import (
	"fmt"
)

func main() {
	channel1 := make(chan string)
	channel1 = nil
	channel2 := make(chan string)

	select {
	case message1 := <-channel1:
		fmt.Println(message1)
	case message2 := <-channel2:
		fmt.Println(message2)
	default:
		fmt.Println("No channel is ready")
	}
}

/*
If we want to disable the send/receive operations of a channel in select statement,
we can use nil channels. Let’s add other cases and the default case to resolve the deadlock:
output : No channel is ready
*/
