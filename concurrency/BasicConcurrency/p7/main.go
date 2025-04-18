package main

import "fmt"

func sendValues(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}
	close(myIntChannel)
}

func main() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for value := range myIntChannel {
		fmt.Println(value) //receiving value
	}
}

/*
Now we got exactly the same results but we don’t have to manually
check if the channel is open or not. This was just syntactic sugar.
*/
