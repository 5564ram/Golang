package main

import (
	"fmt"
)

var resume chan int
var count = 0

func integers() chan int {
	yield := make(chan int)

	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

func main() {
	resume = integers()
	fmt.Println(generateInteger()) //=> 0
	fmt.Println(generateInteger()) //=> 1
	fmt.Println(generateInteger()) //=> 2
	fmt.Println("count: ", count)
}
