package main

import "fmt"

func sendValues(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}

}

func main() {
	myIntChannel := make(chan int)
	defer close(myIntChannel)
	go sendValues(myIntChannel)

	for i := 0; i < 5; i++ {
		fmt.Println(<-myIntChannel) //receiving value
	}
	// the below code will give deadlock as it will be waiting to receive the 6th value as channel is not closed.
	// for value := range myIntChannel {
	// 	fmt.Println(value) //receiving value
	// }
}

/*In general, it is good practice to defer the closing of channels in the main program
so that we clean up everything ourselves.
*/
