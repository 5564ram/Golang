package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Race(channel, quit chan string, i int) {

	channel <- fmt.Sprintf("Car %d started!", i)
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		quit <- fmt.Sprintf("Car %d reached the finishing line!", i)
		fmt.Println(<-quit)
		wg.Done()
	}
}

func main() {

	channel := make(chan string)
	quit := make(chan string)
	wg.Add(1)
	for i := 0; i < 3; i++ {
		go Race(channel, quit, i)
	}

	for {
		select {
		case raceUpdates := <-channel:
			fmt.Println(raceUpdates)
		case winnerAnnoucement := <-quit:
			fmt.Println(winnerAnnoucement)
			quit <- "You win!"
			wg.Wait()
			return

		}
	}
}

/*
First of all, we declare a WaitGroup named wg on line 7. The purpose is to send data back to the
winning routine and halting the program once a routine wins.

Look at line 25. We set the counter of wg to 1 by wg.Add(1).
This means we’ll only wait for one routine. As you can see on line 36, we send back
the winning message to the goroutine that won, i.e., which sent the winning message to the quit channel
first. In response, the quit channel also communicates to the goroutine that they have won (line 16).

At line 17, we add wg.Done() which decrements the counter once invoked. Now the counter is 0.
As soon as the counter reaches 0, the program moves on to the next statement after wg.Wait()
(at line 37) which is return in this case and therefore, the program exits.
*/
