package main

import (
	"fmt"
)

const TERM = 5

func main() {
	i := 0
	c := make(chan int)

	go terms(TERM, c)
	for {
		if result, ok := <-c; ok {
			fmt.Printf("value(%d) is: %d\n", i, result)
			i++
		}
	}
}

func terms(TERM int, c chan int) {
	for i := 0; i <= TERM; i++ {
		c <- calculate(i)
	}
	close(c)
}

func calculate(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = calculate(n-1) + calculate(n-2)
	}
	return
}
