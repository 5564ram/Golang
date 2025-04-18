package main

import "fmt"

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 3:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add2() func(b int) int { // return a function
	return func(b int) int {
		return b + 2
	}
}
func Adder(a int) func(b int) int { // return a function
	return func(b int) int {
		return a + b
	}
}

/*
--------- Returning a function using closures ----------

In the above program, we implement two functions Add2 and Adder, which return another lambda function.
There is a header of Add2 at line 13. You can notice that it returns an anonymous function that takes
b (of type int) as a parameter and returns an int value. That anonymous function is returning b+2.
Similarly, there is a header Adder at line 18. You can notice that it takes a parameter a (of type int),
and it also returns a closure that takes b (of type int) as a parameter and returns an int value.
That closure is returning a+b.

Now, look at line 6 in main. We are calling Add2 and setting it equal to p2. Then, we call p2(3) at line 7,
which calls the closure returned by Add2. So b of Add2 is equal to 3. On returning b+2, 5 will be printed.
Similarly, look at line 9 in main. We are calling Adder(2) and setting it equal to TwoAdder. Then, we call
TwoAdder(3) at line 10, which calls the closure returned by Adder. So b of Adder is equal to 3, and a of
Adder is equal to 2. On returning a+b, 5 will be printed.
*/
