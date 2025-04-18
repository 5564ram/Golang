package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	// fmt.Println(f(), f(), f(), f(), f())
	for i := 0; i <= 9; i++ {
		fmt.Println(i+2, f())
	}
}

/*

In the program above, we implement function fib, which returns another lambda function.
In the header of fib at line 6, notice that it returns an anonymous function that takes nothing as a parameter
and returns an int value. We declare an int variable a and b in fib at line 7 and initialize both of them with 1.
That anonymous function returns the modified value of b, after adding a into b. Now, look at line 15 in main.
The function fib() is now assigned to the variable f (which is then of type func() int). At line 18, there is a
for loop that iterates 10 times. In each iteration, we call f(). This means the first 10 Fibonacci values will
be printed, excluding values of 0 and 1. Let’s follow the first 5 iterations.

In the first iteration, a and b are declared and initialized with 1. Now, the lambda function makes b equal
to a+b(which is 2) and a equal to b (which is 1). The value 2 will be returned as b holds the Fibonacci value
at a position and a holds previous Fibonacci value.

In the second iteration, a and b will hold their values from the previous iteration.
Now, the function makes b equal to a+b(which is 3) and a equal to b (which is 2).
The value b will be returned which is 3.

In the third iteration, the function makes b equal to a+b(which is 5) and a equal to b (which is 3).
The value b will be returned which is 5.

In the fourth iteration, the function makes b equal to a+b(which is 8) and a equal to b (which is 5).
The value b will be returned which is 8.

In the fifth iteration, the function makes b equal to a+b(which is 13) and a equal to b (which is 8).
The value b will be returned which is 13.
*/
