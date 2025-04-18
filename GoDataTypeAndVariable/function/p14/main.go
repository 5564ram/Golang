package main

import "fmt"

func main() {
	Print(1, 3, 2, 0)
}

func Print(a ...int) { // variable number of parameters
	if len(a) == 0 {
		return
	}

	for _, v := range a {
		fmt.Printf("The number is: %d\n", v)
	}
	return
}

// Passing a variable number of parameters

/*
As you can see, we implemented a function at line 8 called Print that takes the variable number
of parameters to get printed. Now, there can be a case where the user doesn’t pass any parameter.
For this, we simply return from the function at line 10. If the parameters list is not empty,
we print the parameters using the range function at line 14.
*/
