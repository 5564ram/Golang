package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Race(channel, quit chan string, i int) {

	channel <- fmt.Sprintf("Car %d started!", i)
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		quit <- fmt.Sprintf("Car %d reached the finishing line!", i)
	}

}

func main() {

	channel := make(chan string)
	quit := make(chan string)

	for i := 0; i < 3; i++ {
		go Race(channel, quit, i)
	}

	for {
		select {
		case raceUpdates := <-channel:
			fmt.Println(raceUpdates)
		case winnerAnnoucement := <-quit:
			fmt.Println(winnerAnnoucement)
			return
		}
	}
}

/*
After the goroutines start concurrently, we have a select statement from line 27 to line 34.

for{
    select{
      case raceUpdates := <-channel:
        fmt.Println(raceUpdates)
      case winnerAnnoucement := <-quit:
        fmt.Println(winnerAnnoucement)
       return

    }
  }
The channel will update us when the car starts off. If any one of the cars reaches the finishing line,
it means that a goroutine has reached this line:

quit <- fmt.Sprintf("Car %d reached the finishing line!", i)

Then the second case of the select statement is executed and we call return to break out of the for loop,
which is supposed to be infinite.

This pattern is useful when we are supposed to break out of channel communication operations.
*/
