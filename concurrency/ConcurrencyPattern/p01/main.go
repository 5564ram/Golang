package main

import (
	"fmt"
)

func foo() <-chan string {
	mychannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			mychannel <- fmt.Sprintf("%s %d", "Counter at : ", i)
		}
	}()

	return mychannel // returns the channel as returning argument
}

func main() {
	mychannel := foo() // foo() returns a channel.

	for i := 0; i < 10; i++ {
		fmt.Printf("%q\n", <-mychannel)
	}

	fmt.Println("Done with Counter")
}

/*
Generator Pattern ...
Generators return the next value in a sequence each time they are called.
This means that each value is available as an output before the generator
computes the next value. Hence, this pattern is used to introduce parallelism in our program.
In the code above, we use a generator as a function which returns a channel. The foo() function,
when invoked on line 20, calls a goroutine which sends a string onto the channel on line 12.
The string contains the value of i which is incremented in every iteration before being sent
to the channel. As the goroutine executes concurrently, the program returns mychannel immediately
to the main routine. Then mychannel receives the data on the channel one by one on line 23 as it is
being sent by the goroutine. In conclusion, both the goroutine and the main routine can execute concurrently
as we print the value of the counter as soon as we receive it while the goroutine simultaneously
computes the next value.
*/
