package main

import "fmt"

func main() {
	mychannel := make(chan int, 2)

	fmt.Println(<-mychannel)
}

/*
Also, note that the receive operation will block the code if it finds the buffer
to be empty. Check out the code given below which also generates an error
pointing to a deadlock:
In conclusion, we can say that if there are no receive operations for a channel,
a goroutine can still perform c sending operations, where c is the capacity of
the buffered channel. So you can see that buffered channels remove synchronization.
They work in a way similar to mailboxes and their usage depends on the type of problem you have to solve.
*/
