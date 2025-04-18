package main

import (
	"fmt"
	"sync"
)

var wg = new(sync.WaitGroup)

func producer() chan int {
	fmt.Println("producer")
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 1000; i++ {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	fmt.Println("producer done")
	return ch
}
func consumer(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}
func main() {
	consumer(producer())
	wg.Wait()
}
