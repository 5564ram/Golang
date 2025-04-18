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
		close(channel1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel2 <- "I'll print every 1s"
			time.Sleep(time.Second * 1)

		}
		close(channel2)
	}()

	for i := 0; i < 5; i++ {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)

		}
	}
}

/*
The select statement will pick up any channel operation that is ready. As you can see from the output,
there are more I'll print every 100ms statements as compared to I'll print every 1s because channel1
is ready after every 100ms whereas channel2 takes 1 second in order to get ready to send/receive messages.
*/
