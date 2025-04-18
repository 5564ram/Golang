package main

import "fmt"

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
	// ...
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking!")
}

func main() {
	b := new(Bird)
	DuckDance(b)
}

/*
In the code above, at line 4, we define an IDuck interface with two methods: Quack() and Walk().
Look at the header of function DuckDance: func DuckDance(duck IDuck) at line 9. It takes the parameter
of the IDuck interface type. It calls the functions Quack() and Walk() 3 times in a row using a for loop.

Type Bird is defined at line 16. We leave out properties here because they are not needed for the example.
However, Bird implements the two methods of the interface IDuck, Quack() (from line 20 to line 22) and Walk()
(from line 24 to line 26).

We then create a variable b of type Bird in main() at line 29. Now, we can call DuckDance on b
(see line 30).
*/
