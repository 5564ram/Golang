package main

import (
	"fmt"
	"os"
	"strconv"
)

func iter(b int, c chan int) {
	for i := 0; i < b; i++ {
		c <- i // put integers from 0 to n-1 on channel c
	}
	close(c)
}

func main() {
	n, _ := strconv.Atoi(os.Args[1]) // line arg. converted to integer
	a := make(chan int)
	b := make(chan int)
	go iter(n, a)
	go iter(n, b)
	for a != nil || b != nil {
		select {
		case x, ok := <-a: // takes int value from channel a
			if ok {
				fmt.Printf("channel a :  %d\n", x)
			} else { // if channel a is closed
				a = nil
			}
		case y, ok := <-b: // takes int value from channel b
			if ok {
				fmt.Printf("channel b : %d\n", y)
			} else { // if channel b is closed
				b = nil
			}
		}
	}
}

// go run main.go 10
