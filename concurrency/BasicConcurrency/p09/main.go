package main

import "fmt"

func main() {
	mychannel := make(chan int)
	mychannel <- 10
	fmt.Println(<-mychannel)
}

/*Now you’ll be fully clear on why the above code didn’t work. This is because of the sending and
receiving operations which are blocking the code. When we wrap one of them in a goroutine
such that they are ready to unblock each other, the program executes successfully.
*/
