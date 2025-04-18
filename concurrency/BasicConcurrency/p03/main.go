package main

import "fmt"

func sendValues(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}

}

func main() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 5; i++ {
		fmt.Println(<-myIntChannel) //receiving value
	}
}

/*Let’s get into what happened above. In a goroutine created on line 14,
the function sendValues was sending values over myIntChannel by using a for-loop.
On the other hand, on line 17, myIntChannel was receiving values and the program was
printing them onto the console. The most important point to note is that both the following
statements were blocking operations:

myIntChannel <- i
<-myIntChannel
Hence, the program when blocked on myIntChannel <- i was unblocked by the <-myIntChannel statement.
This was only possible as they were running concurrently.*/
