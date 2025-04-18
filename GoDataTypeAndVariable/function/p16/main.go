package main

import "fmt"

type flt func(int) bool

// isOdd takes an int slice and returns a bool set to true if the
// int parameter is odd, or false if not.
// isOdd is of type func(int) bool which is what flt is declared to be.

func isOdd(n int) bool {
	return n%2 == 1
}

// Same comment for isEven
func isEven(n int) bool {
	return n%2 == 0
}

func filter(sl []int, f flt) []int {
	var res []int
	for _, val := range sl {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)
}

/*
Function used as a filter
The program has two basic functions. The first function, isOdd takes n as a parameter and returns a
boolean value (see its header at line 9). If n is odd, it will return true. Otherwise, it will return false.
Similarly, the second function, isEven takes n as a parameter and returns a boolean value
(see its header at line 17). If n is even, it will return true. Otherwise, it will return false.

At line 4, we are aliasing a type. A function that takes a single integer as a parameter and returns
a single boolean value is given a type flt. Now, moving towards a major part of the program: the filter
function. See its header at line 24. It takes a slice of integers (that are to be judged as even or odd)
as a first parameter, and function f of type flt (either isEven or isOdd). The function filter returns a
slice of integers res, for which the f returns true.

Let’s see the main function now. At line 35, we declare a slice of integers named slice.
Then at line37, we call the filter function with slice as the first parameter and isOdd as
the second parameter and store the result in odd slice. Similarly, at line39, we call the filter
function with slice as the first parameter and isEven as the second parameter and store the result
in the even slice. Printing odd and even slices at line 38 and line 40, respectively, verifies the
result. The same principle can be applied to construct filters for any type.
*/
