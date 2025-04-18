package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 10)
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}
	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}

/*

Slice at 0 is 0
Slice at 1 is 5
Slice at 2 is 10
Slice at 3 is 15
Slice at 4 is 20
Slice at 5 is 25
Slice at 6 is 30
Slice at 7 is 35
Slice at 8 is 40
Slice at 9 is 45

The length of slice1 is 10
The capacity of slice1 is 10

In the program above, we make a slice via the make function. Look at line 5 in main.
We make an integer type slice slice1 with the length 10. The loop at line 7 is populating
the values in slice1. Each element at index i is given a value of 5*i at line 8. The next
for loop at line 11 is printing the slice elements. Then, at line 14 and line 15, we are
printing the length and capacity of slice1 using the len() and cap() functions, which are both 10.


s2 := make([]int, 10)
the following is true:

cap(s2) == len(s2) == 10
The function make takes 2 parameters:

the type to be created
the number of items in the slice. If you want slice1 not to occupy the whole array (with length cap)
from the start, but only a number len of items, use the form:

slice1 := make([]type, len, cap)
So make has the signature:

func make([]T, len, cap) // []T with type T and optional parameter cap
And the following statements result in the same slice:

make([]int, 50, 100)
new([100]int)[0:50]
*/
