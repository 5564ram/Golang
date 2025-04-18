package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), " , ")
	fmt.Print(f(20), " , ")
	fmt.Print(f(300))
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

/*
In the above program, we implement function Adder which returns another lambda function.
There is a header of Adder at line 11. You can notice that it returns an anonymous function
that takes an int parameter and returns an int value. We declare an int variable x in Adder at
line 12, and that anonymous function returns the modified value of x, after adding its parameter
delta into x.

Now, look at line 5 in main. Adder() is now assigned to the variable f (which is then of type func(int) int).
In the calls to f, delta in Adder() gets the values 1, 20 and 300 from line 6, line 7 and line 8, respectively.
We see that between the calls of f the value of x is retained, first it is: 0 + 1 = 1, then it becomes 1 + 20 = 21,
then 21 is added to 300 to give the result 321. The lambda function stores and accumulates the values of its
variables. It still has access to the (local) variables defined in the current function.
*/
