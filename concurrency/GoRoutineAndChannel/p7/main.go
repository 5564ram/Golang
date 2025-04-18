package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go sendData(ch, wg)
	go getData(ch, wg)
	wg.Wait()
}

func sendData(ch chan string, wg *sync.WaitGroup) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
	defer func() {
		wg.Done()
		close(ch)
	}()

}

func getData(ch chan string, wg *sync.WaitGroup) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
	fmt.Println()
	defer wg.Done()
}
