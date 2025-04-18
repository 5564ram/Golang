package main

import "fmt"

func deposit(balance *int, amount int) {
	*balance += amount //add amount to balance
}

func withdraw(balance *int, amount int) {
	*balance -= amount //subtract amount from balance
}

func main() {

	balance := 100

	go deposit(&balance, 10) //depositing 10

	withdraw(&balance, 50) //withdrawing 50

	fmt.Println(balance)

}

/* A race condition is a flaw in a program regarding the timing/ordering of operations which disrupts the logic of the program and leads to erroneous results.


In the code above, we have a balance of 100. We first execute the deposit() operation using a goroutine and deposit 10 to our balance.
Next, we withdraw 50 from our balance which makes the final balance 60.
However, if you run the code, the final value of balance will come out to be 50.
This is as a result of a race condition which has compromised the correctness of the program due to an incorrect order of execution of operations.*/

// go run -race main.go
