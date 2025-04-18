package main

import "fmt"

func sendValues(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}

}

func main() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 6; i++ {
		fmt.Println(<-myIntChannel) //receiving value
	}
}

/*
Finally, we reached a deadlock. Any idea why it happened? Let’s figure it out.

So I just changed the for loop condition in the main loop from i < 5 to i < 6.
As a result, the main routine is blocked on <-myIntChannel because the sending
operation has sent only 5 values which were received by the 5 iterations of the loop.
However, for the 6th iteration, there is no sending operation that will send value on
the channel. Therefore, the program is blocked on the receiving operation resulting in
a deadlock.

One way to fix this problem is by closing the channel.
*/
