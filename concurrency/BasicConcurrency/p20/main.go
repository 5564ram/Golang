package main

import (
	"fmt"
	"sync"
)

func WelcomeMessage() {
	fmt.Println("Welcome to Educative!")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		WelcomeMessage()
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello World!")
		wg.Done()
	}()

	wg.Wait()
}

/*
A WaitGroup blocks a program and waits for a set of goroutines to finish before
moving to the next steps of execution.

We can use WaitGroups through the following functions:

.Add(int): This function takes in an integer value which is essentially the number of goroutines
which the waitgroup has to wait for. This function must be called before we execute a goroutine.

.Done(): This function is called within the goroutine to signal that the goroutine has successfully executed.

.Wait(): This function blocks the program until all the goroutines specified by Add() have invoked Done()
 from within.

 First of all, we import the sync package and declare a WaitGroup named wg on line 10.
 As we have two goroutines that have to be finished before we exit our program, we set
 the counter of wg to 2 by wg.Add(2) on line 11. Next, we write wg.Done() at the end of
 the execution of the goroutines (line 14 and line 18). This is to give a signal that the
 goroutine has finished. Finally, we add wg.Wait() at the very end of our main program on
 line 21. As soon as the execution jumps to wg.Wait(), the program will halt until we get
 the signals from the goroutines by wg.Done(). wg.Done() decrements the counter once invoked.
 As soon as the counter reaches 0, the program moves on to the next statement after wg.Wait()
 which is none in this case and therefore, the program exits.

*/
