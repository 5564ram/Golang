package main

import (
	"fmt"
	"sync"
)

func printTable(n int, wg *sync.WaitGroup) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, n*i)
		// time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for number := 2; number <= 12; number++ {
		wg.Add(1)
		go printTable(number, &wg)
		// wg.Wait() // it is not cucurrent execution type as it is sequential
	}
	wg.Wait()

}

/*
So we added a WaitGroup from the sync package on line 15 named wg.
At every iteration in the for-loop in the main routine, we increment
the counter of wg on line 18. In the next step on line 19, we execute
the printTable function in a goroutine and pass number and wg as input
arguments. After launching all the goroutines for number = 2 to number = 12
in the for-loop which sets the counter of wg to 11, we proceed to line 22 and
call Wait() on wg. This blocks the main routine until the counter of wg equals 0.
Hence, our main routine cannot exit until and unless we are done with the printing
inside the printTable functions.
*/
