package main

import (
	"fmt"
	"sync"
)

func HeavyFunction1(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do a lot of stuff
}

func HeavyFunction2(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do a lot of stuff
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go HeavyFunction1(wg)
	go HeavyFunction2(wg)
	wg.Wait()
	fmt.Println("All Finished!")
}

// This can be accomplished with a coordinating channel: the semaphore pattern, which we’ll study later in this chapter.
//  Another solution is to use the type sync.WaitGroup, which exactly serves this purpose: a WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add on the WaitGroup object to set the number of goroutines to wait for. Every goroutine takes a pointer to the WaitGroup value as a parameter when it is invoked. When each of the goroutines runs, it calls Done when finished.
// The main() calls the Wait() method to block itself until all the goroutines have finished. The following code snippet illustrates this:

// To get a synchronized behavior that is not dependent on time, a function launched in a goroutine can register to a Waitgroup instance, here wg, as one of its parameters. At line 18, we construct the WaitGroup wg, and at line 19, we say that it has to wait for two goroutines.

// Line 20 and line 21 start our goroutines, passing wg as a parameter. Both goroutines will signal wg just before their end of execution with wg.Done() (see line 8 and line 13). Then, the WaitGroup waits in the main() function at line 22, until it has received as many Done signals as there were goroutines registered. Only after that, processing continues with line 23.
