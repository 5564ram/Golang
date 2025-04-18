package main

import (
	"fmt"
	"sync"
)

func generate(wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i < 100; i++ {
			out <- i
		}
		close(out) // Close channel to signal end of numbers
	}()
	return out
}

func filter(in chan int, prime int, wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(out) // Close output channel when done
		}()

		for num := range in { // Using range prevents reading from a closed channel
			if num%prime != 0 {
				out <- num
			}
		}

	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	wg := new(sync.WaitGroup) // WaitGroup to track goroutines

	go func() {
		defer close(out) // Close `out` once all primes are found
		in := generate(wg)
		for {
			prime, ok := <-in
			if !ok { // If channel is closed, break the loop
				// wg.Wait() // Wait for all goroutines to finish before exiting
				return
			}
			out <- prime
			in = filter(in, prime, wg)
		}
	}()
	return out
}

func main() {
	primes := sieve()
	for prime := range primes {
		fmt.Println(prime)
	}
}
