package main

import (
	"fmt"
)

func main() {
	channel1 := make(chan string)
	channel1 = nil

	select {
	case message1 := <-channel1:
		fmt.Println(message1)
	}
}

/*
Nil channel will block the select statement forever, giving an error stating a deadlock:
*/
