package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dynamite := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		dynamite <- "Dynamite Diffused!"
	}()

	timeout := time.After(time.Duration(rand.Intn(500)) * time.Millisecond)
	for {
		select {
		case s := <-dynamite:
			fmt.Println(s)
			return
		case <-timeout:
			fmt.Println("Dynamite Explodes!")
			return
		}
	}
}

/*
n the above code example, we have a channel named dynamite. On line 10, we create a goroutine which
sends Dynamite Diffused! on the dynamite channel after a random period of time.

go func(){
    rand.Seed(time.Now().UnixNano())
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    dynamite <- "Dynamite Diffused!"
}()

From line 17 to line 24, we have a select statement comprising of two cases.

select
{
  case s := <-dynamite:
    fmt.Println(s)
   case <-time.After(time.Duration(rand.Intn(500)) * time.Millisecond):
    fmt.Println("Dynamite Explodes!")
}

If the dynamite channel is able to deliver the Dynamite Diffused message before the random period
of time that will execute the second case, we are safe! But if the time.After function delivers the
current time after a random duration, the bomb will explode and we are doomed!
*/
