package main

import "fmt"

func main() {
	mychannel := make(chan int, 2)
	mychannel <- 10
	fmt.Println(<-mychannel)
}

/*
As you can see, we don’t get a deadlock as before.
This is because the send operation on line 6 does not block the code
until we have reached our capacity.
*/
