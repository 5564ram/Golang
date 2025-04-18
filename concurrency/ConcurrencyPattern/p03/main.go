package main

import "fmt"

func fibonacci(n int) chan int {
	mychannel := make(chan int)
	go func() {
		i, j := 0, 1
		for k := 0; k < n; k++ {
			mychannel <- i
			i, j = i+j, i
		}
		close(mychannel)
	}()
	return mychannel
}

func main() {
	val := fibonacci(10)
	for i := range val {
		//do anything with the nth term while the fibonacci()
		//is computing the next term
		fmt.Println(i)
	}
}

/*
The crux of the solution lies in the part where we return a channel from the function on line 16
to the main routine on line 21 which allows the fibonacci() function to communicate each and every
term to the main routine as soon as it is computed in the function body. In this way, when we print
a term in the Fibonacci sequence, this pattern enables us to do whatever we want to do with each term
while it is concurrently computing the next term.
*/
