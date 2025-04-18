package main

import "fmt"

func sendValues(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}
	close(myIntChannel)
}

func main() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 6; i++ {
		fmt.Println(<-myIntChannel) //receiving value
	}
}

/*
You can see that when we close the channel after all our send operations,
the receive operation returns 0 without blocking on the 6th iteration.

Additionally, the receive operation returns another value with the data to
indicate whether the channel is open or not. Let’s see how we can use it to solve
our problem:
*/
