package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(4 * 1e9) // sleep works with a Duration in nanoseconds (ns) !
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(1 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}

// The 3 functions main(), longWait() and shortWait() are started in this order as independent processing units, and then work concurrently. Each function outputs a message at its beginning and at the end of its processing. To simulate their processing times, we use the Sleep function from the time package. Sleep() pauses the processing of the function or goroutine for the indicated amount of time, which is given in nanoseconds (ns, the notation 1e9 represents 1 times 10 to the 9th power, e = exponent). They print their messages in the order we expect, always the same. But, we see clearly that they work simultaneously, in parallel. We let the main() pause for 10s, so we are sure that it will terminate after the two goroutines. If not (say if we let main() stop for only 4s), main() stops the execution earlier and longWait() doesn’t get the chance to complete. If we do not wait in main(), the program stops and the goroutines die with it.

// When the function main() returns, the program exits. It does not wait for other (non-main) goroutines to complete. That is the reason why in server-programs where each request is handled by a response started as a goroutine, the server() function, which starts the goroutines, must be kept alive. This is usually done by starting it as an infinite loop.

// Moreover, goroutines are independent units of execution, and when a number of them start one after the other, you cannot depend on when a goroutine will actually be started. The logic of your code must be independent of the order in which goroutines are invoked.
