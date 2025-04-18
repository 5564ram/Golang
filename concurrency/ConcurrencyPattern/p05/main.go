package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var myNumbers [10]int
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		myNumbers[i] = rand.Intn(50)
	}
	fmt.Println("My numbers are: ", myNumbers)

	mychannelOut := channelGenerator(myNumbers)

	mychannel1 := double(mychannelOut)
	mychannel2 := double(mychannelOut)

	mychannelIn := fanIn(mychannel1, mychannel2)

	for i := 0; i < len(myNumbers); i++ {
		fmt.Println(<-mychannelIn)
	}
}

func channelGenerator(numbers [10]int) <-chan int {
	channel := make(chan int)
	go func() {
		for _, i := range numbers {
			channel <- i
		}
		close(channel)
	}()
	return channel
}

func double(inputchannel <-chan int) <-chan string {
	channel := make(chan string)
	go func() {
		for num := range inputchannel {
			channel <- fmt.Sprintf("%d * 2 = %d", num, num*2)
		}
		close(channel)
	}()
	return channel
}

func fanIn(inputchannel1, inputchannel2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			select {
			case message1 := <-inputchannel1:
				channel <- message1
			case message2 := <-inputchannel2:
				channel <- message2
			}
		}
	}()
	return channel
}

/*
Let’s analyze what’s happening above. From line 12 to line 15, we are just populating an array with
random numbers ranging from 0 to 49.

for i := 0; i < 10; i++{
        rand.Seed(time.Now().UnixNano())
        myNumbers[i]=rand.Intn(50)
}
On line 17, we create a common channel mychannelOut using the channelGenerator function.

The channelGenerator function is as follows:

func channelGenerator(numbers [10]int) <-chan string {
    channel := make(chan string)
    go func() {
        for _, i := range numbers {
            channel  <-  strconv.Itoa(i)
        }
        close(channel)
    }()
    return channel
}
It returns a channel which will receive data from the goroutine created in the function itself.
We are just converting the integer values from the input array to string and sending on to the channel.

In the main routine (lines 19-20), we fan-out our common channel mychannel to two goroutines, which are
created from inside the function double.

func double(inputchannel <-chan string) <-chan string {
    channel := make(chan string)
    go func() {
        for i := range inputchannel {
            num, err := strconv.Atoi(i)
             if err != nil {
                // handle error
             }
             channel <- fmt.Sprintf("%d * 2 = %d", num,num*2)
        }
        close(channel)
    }()
    return channel
}
The double function takes our common channel as an input and sends data on to another channel.
Note that there are two instances of the double function, which implies that the two goroutines are
concurrently receiving data from a single channel. This is pretty much our fan-out technique where we
distribute the data sent from one channel to two channels concurrently, which divides the computation
of doubling the integers and returns them as strings among the two goroutines.

On line 22, we fan-in the results from the channels returned by the double function and print them on
to the console.
*/
