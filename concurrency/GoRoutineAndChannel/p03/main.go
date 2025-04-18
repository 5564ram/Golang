package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch) // calling goroutine to send the data
	go getData(ch)  // calling goroutine to receive the data
	time.Sleep(1e9)
}

func sendData(ch chan string) { // sending data to ch channel
	fmt.Println("Sending data Washington ", len(ch))
	ch <- "Washington"
	fmt.Println("Sending data Tripoli ", len(ch))
	ch <- "Tripoli"
	fmt.Println("Sending data London ", len(ch))
	ch <- "London"
	fmt.Println("Sending data Beijing ", len(ch))
	ch <- "Beijing"
	fmt.Println("Sending data Tokyo ", len(ch))
	ch <- "Tokyo"
	defer close(ch) // closed the channel
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch // receiving data sent to ch channel
		fmt.Println("Received data: ", input)

		// fmt.Printf("%s ", input)
	}

}

/* In main() at line 8, a channel ch of strings is made. Then, two goroutines are started: sendData() (at line 9) sends 5 strings over channel ch
(see implementation from line 14 to line 20), and getData() (at line 10) receives them one by one in order in the string input (line 25) and prints what is received (line 26).
. Experiment what happens when you comment out time.Sleep(1e9). Also, remember to comment out the import of the time package or your code won’t compile.

Here, we see that synchronization between the goroutines becomes important:

The main() waits for 1 second so that both goroutines can come to completion. If this is not allowed, sendData() doesn’t have the chance to produce its output.
getData() works with an infinite for-loop. This comes to an end when sendData() has finished, and ch is empty.
If we remove one or all go keywords, the program doesn’t work anymore, and the Go runtime throws a panic:
---- Error run <path> with code Crashed Fatal error: all goroutines are asleep - deadlock!

Why does this occur? The runtime can detect that all goroutines (or perhaps only one in this case) are waiting
for something (to write to a channel or be able to read from a channel), which means the program can’t possibly proceed.
 It is a form of a deadlock, and the runtime can detect it for us.*/
