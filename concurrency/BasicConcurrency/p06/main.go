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

	for i := 0; i < 6; i++ {
		value, open := <-myIntChannel
		if !open {
			break
		}
		fmt.Println(value) //receiving value
	}
}

/*
Here, we check if the channel is open or not using open (line 18) and break the loop
if the channel is closed (line 19).

Another way to implement the same functionality as above is by using the range function
on line 16:
*/
