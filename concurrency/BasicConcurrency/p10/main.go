package main

import "fmt"

func main() {
	mychannel := make(chan int)
	go func() {
		mychannel <- 10
	}()
	fmt.Println(<-mychannel)
}

/*
we wrap one of them in a goroutine such that they are ready to unblock each other, the program executes successfully.
*/
