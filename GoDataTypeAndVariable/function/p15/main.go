package main

import (
	"fmt"
)

func main() {
	callback(1, Add) // function passed as a parameter
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

/*
Function used as a parameter
To understand the code, look at line 10 at the function header of the function Add.
It takes two parameters a and b and prints the sum of the parameters. In the main,
we are just calling the callback function as: callback(1, Add). Here, we have
two parameters: the first is 1 used as int and the second is Add function passed as a
parameter. You can verify it through the header of the callback function: func
callback(y int, f func(int, int)) at line 14. At line 15, we call the function f from
the header of callback, and treat it like the Add function. So y in callback is a in Add,
and 2 is b in Add. So, the output will be: The sum of 1 and 2 is: 3.
*/
