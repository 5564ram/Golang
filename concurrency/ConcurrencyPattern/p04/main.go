package main

import (
	"fmt"
	"math/rand"
	"time"
)

func updatePosition(name string) <-chan string {
	positionChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			positionChannel <- fmt.Sprintf("%s %d", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return positionChannel
}

func fanIn(mychannel1, mychannel2 <-chan string) <-chan string {
	mychannel := make(chan string)

	go func() {
		select {
		case message1 := <-mychannel1:
			mychannel <- message1
		case message2 := <-mychannel2:
			mychannel <- message2
		}
	}()
	return mychannel
}

func main() {
	positionsChannel := fanIn(updatePosition("Legolas :"), updatePosition("Gandalf :"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-positionsChannel)
	}
	fmt.Println("Done with getting updates on positions.")
}

/*

------ Fan-In, Fan-Out ------
Fan-In refers to a technique in which you join data from multiple inputs into a single entity.
On the other hand, Fan-Out means to divide the data from a single source into multiple smaller
chunks. In this lesson, we’ll learn how to make use of both these techniques.

By using Fan-In technique, we’ll combine the inputs from both channels and send them through a single channel.
Look at the code below to see how it’s done:

The crux of our technique lies in the fanIn function:
func fanIn(mychannel1, mychannel2 <-chan string) <-chan string {
    mychannel := make(chan string)

    go func() {
        for {
            mychannel <- <-mychannel1
        }
    }()

    go func() {
        for {
            mychannel <- <-mychannel2
        }
    }()

    return mychannel
}
Here we take in two channels as input and specify one channel as the return argument,
i.e., mychannel which is the Fan-In channel. Afterward, we declare two goroutines from
line 23 to line 33 which receive data from mychannel1 and mychannel2 and send it to our Fan-In channel,
mychannel on line 25 and on line 31. Hence, data will be passed to mychannel as soon as it is received
by mychannel1 and mychannel2 because the goroutines are running concurrently.

Additionally, we pass the updatePosition function for both Legolas and Gandalf into our fanIn function on
line 40. The rest of the logic is the same as before. The only difference comes from the fact that mychannel1
and mychannel2 are communicating with mychannel now instead of directly communicating with the main routine
as before. You will realize from the output that the position updates are no longer sequential. Thus, by using
this technique, we can solve the blocking issue that we were previously facing.
*/
