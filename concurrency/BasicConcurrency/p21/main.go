package main

import (
	"fmt"
	"sync"
)

func WelcomeMessage(wg *sync.WaitGroup) {
	fmt.Println("Welcome to Educative!")
	wg.Done()
}

func main() {

	// wg := new(sync.WaitGroup)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go WelcomeMessage(wg)
	go func() {
		fmt.Println("Hello World!")
		wg.Done()
	}()

	wg.Wait()

}

/*
Note: A WaitGroup must not be copied after first use.

This is because it can affect our counter which will disrupt the logic of our program.

In summary, waitgroups provide us with a way to make sure that all concurrent operations
reach completion before the program exits. They can be vital in applications where you
have to block code. However, they can only be used if we don’t want any output from those operations.
*/
