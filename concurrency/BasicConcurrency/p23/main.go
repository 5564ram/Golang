package main

import (
	"fmt"
	"sync"
)

func deposit(balance *int, amount int, myMutex *sync.Mutex, myWaitGroup *sync.WaitGroup) {
	myMutex.Lock()
	*balance += amount //add amount to balance
	myMutex.Unlock()
	myWaitGroup.Done()
}

func withdraw(balance *int, amount int, myMutex *sync.Mutex, myWaitGroup *sync.WaitGroup) {
	myMutex.Lock()
	*balance -= amount //subtract amount from balance
	myMutex.Unlock()
	myWaitGroup.Done()
}

func main() {

	balance := 100
	var myWaitGroup sync.WaitGroup
	var myMutex sync.Mutex

	myWaitGroup.Add(2)
	go deposit(&balance, 10, &myMutex, &myWaitGroup) //depositing 10
	withdraw(&balance, 50, &myMutex, &myWaitGroup)   //withdrawing 50

	myWaitGroup.Wait()
	fmt.Println(balance)

}

/*
The variable balance is a shared resource of the deposit() and withdraw functions.
Therefore, whenever we encounter any operations that access and update balance such as:
*balance += amount and *balance -= amount
we surround that code with our locking mechanism. Once a lock has been acquired by a goroutine
by myMutex.Lock(), no other goroutine is allowed to enter the critical section until and unless
the critical section is executed and the lock is released through myMutex.Unlock().
*/
