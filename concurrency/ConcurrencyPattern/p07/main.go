package main

import "fmt"

type Money struct {
	amount int
	year   int
}

func sendMoney(parent chan Money) {

	for i := 0; i <= 18; i++ {
		parent <- Money{5000, i}
	}
	close(parent)
}

func main() {
	money := make(chan Money)

	go sendMoney(money)

	for kidMoney := range money {
		fmt.Printf("Money received by kid in year %d : %d\n", kidMoney.year, kidMoney.amount)
	}
}

/*
In Go, we have a range function which lets us iterate over elements in different data structures.
Using this function, we can range over the items we receive on a channel until it is closed.
Also, note that only the sender, not the receiver, should close the channel when it feels that it
has no more values to send.
This is because the receiver may still be waiting for values to be sent on the channel.
The receiver should never close the channel as it may lead to a panic in the program.
The sender should close the channel when it is done sending values to the channel.
The receiver should only read the values from the channel and not close it.
The code above demonstrates how to send a struct on a channel.
The sendMoney function sends a struct of type Money on the channel parent.
The struct contains two fields, amount and year.
The amount is set to 5000 and the year is set to the value of i in the for loop.
The channel is closed after sending all the values.
The main function creates a channel of type Money and calls the sendMoney function in a goroutine.
The main function then ranges over the channel and prints the values received on the channel.
The range function will keep receiving values from the channel until it is closed.
The main function will print the values received on the channel until the channel is closed.
*/
