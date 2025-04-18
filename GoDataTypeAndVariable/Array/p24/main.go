package main

import "fmt"

func sum(a []int) int { // function that sums integers
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var arr = [5]int{0, 1, 2, 3, 4} // declare an array
	fmt.Println(sum(arr[:]))        // passing slice to the function
}

/*
In the above program, we have a function sum, that takes slice of an array a as a parameter and
returns the sum of the elements present in a. Look at its header at line 4. In the main function,
at line 13, we declare an array arr of length 5 and pass it to the sum function on the next line.
We write arr[:], which is enough to pass the array arr as the reference.
*/
