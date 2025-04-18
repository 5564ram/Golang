package main

import (
	"fmt"
)

func tel(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func main() {
	ok := true
	ch := make(chan int)

	go tel(ch)
	for ok {
		i, ok := <-ch
		fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
	}
}
