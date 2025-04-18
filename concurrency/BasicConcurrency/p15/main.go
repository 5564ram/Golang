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

		fmt.Println(<-channel1)
		fmt.Println(<-channel2)

	}
}

/*
Let’s try to understand the usage of a select statement with a simple example consisting
of two channels which are communicating using send/receive operations:

The first goroutine sends a message I'll print every 100ms and then waits for 100ms before sending the message
again in the next iteration. The second goroutine sends a message I'll print every 1s and then waits
for 1s before sending the message again in the next iteration.

On the other hand, the second goroutine functions similarly and sends a message I'll print every 1s in each
of the five iterations with a delay of 1s.

fmt.Println(<-channel1) receives and prints the message I'll print every 100ms and the program moves to fmt.
Println(<-channel2). After receiving and printing the message from the second goroutine, we get done with our
first iteration. However, in the second iteration, after channel1 receives the message sent, our program is
blocked by the second goroutine as it is waiting for 1 second of the previous iteration. Note that channel1
is always ready to send messages but is blocked by the delay in the second goroutine. channel1 has to wait
for 1 second for fmt.Println(<-channel2) to execute every time although the first goroutine may have finished
entirely in less than a second!
*/
