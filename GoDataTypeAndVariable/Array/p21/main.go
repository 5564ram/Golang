package main

import "fmt"

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // dereferencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}

/*
var arrAge = [5]int{18, 20, 15, 22, 16}
var arr = [10]int{ 1, 2, 3 }
This is an array of 10 elements with the 1st three different from 0.

var arrLazy = [...]int{5, 6, 7, 8, 22}
... indicates the compiler has to count the number of items to obtain the length of the array

The index: value syntax can be followed. For example:
var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}

In the above code, there is a function Sum that takes an array and calculates the sum of its elements and returns their sum.
Look at its header at line 11: func Sum(a *[3]float64) (sum float64). The function is taking the pointer to the array a of length 3,
whose elements are of type float64, and returning a float64 number sum. Now, look at the main, where we declared an array called
array as: array := [3]float64{7.0, 8.5, 9.1} at line 5. In the next line, we call the Sum function and pass &array because Sum accepts
a pointer to the array and stores the result in x. At line 8, we are printing x to verify the result.

*/
