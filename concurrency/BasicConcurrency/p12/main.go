package main

import "fmt"

func main() {
	mychannel := make(chan int, 2)
	mychannel <- 10
	mychannel <- 20
	mychannel <- 30
	fmt.Println(<-mychannel)
}

/*
The code above generates an error as mychannel has a capacity of 2 and we already
sent the data two times over the channel. Hence mychannel <- 30 on line 8 blocks
the code resulting in a deadlock because it is our third attempt to send data over
the channel which has a capacity of 2.
*/
