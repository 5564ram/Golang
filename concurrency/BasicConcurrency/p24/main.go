package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	myMutex := sync.Mutex{}
	myMutex.Lock()
	go func() {
		myMutex.Lock()
		fmt.Println("I am in the goroutine")
		myMutex.Unlock()
	}()
	fmt.Println("I am in main routine")
	myMutex.Unlock()
	time.Sleep(time.Second * 1)

}

/*
So first we acquire the lock on line 9 and then we cannot hold another lock on line 11.
The lock is released on line 16 after printing I am in main routine on to the console.
Only then, the goroutine is able to acquire the lock on line 11 and print I am in the goroutine
before releasing that lock on line 13 and exiting the program.

RWMutex stands for Reader/Writer mutual exclusion and is essentially the same as Mutex,
but it gives the lock to more than one reading process or just a writing process.

RWMutex provides us with more control over memory. It allows multiple readers to read the data
at the same time, but only one writer can write the data at a time. This is useful when we have
a large number of goroutines reading the data and only a few goroutines writing the data.
*/
