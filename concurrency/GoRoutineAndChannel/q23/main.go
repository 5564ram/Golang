package main

import (
	"fmt"
	"os"
)

func tel(ch chan int, quit chan bool) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	quit <- true
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go tel(ch, quit)
	for {
		select {
		case i := <-ch:
			fmt.Printf("The counter is at %d\n", i)
		case <-quit:
			os.Exit(0)
			// break : gives a  deadlock on select (break only exits from select, not from for!)
		}
	}
}
