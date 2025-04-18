package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1] // reslice
		slice1[i] = i

		fmt.Printf("The length of slice is %d\n", len(slice1))
	}
	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}

/*

The length of slice is 1
The length of slice is 2
The length of slice is 3
The length of slice is 4
The length of slice is 5
The length of slice is 6
The length of slice is 7
The length of slice is 8
The length of slice is 9
The length of slice is 10
Slice at 0 is 0
Slice at 1 is 1
Slice at 2 is 2
Slice at 3 is 3
Slice at 4 is 4
Slice at 5 is 5
Slice at 6 is 6
Slice at 7 is 7
Slice at 8 is 8
Slice at 9 is 9

In the code above, at line 5 in main, we make a slice slice1 via the make function.
Its length is 0 and its capacity is 10. At line 7 is the for loop that is iterating capacity times.
At line 8, we are reslicing slice1 as: slice1 = slice1[0:i+1]. The length of slice1 is initially 0.
However, reslicing makes the length of slice1 i+1, in every iteration. So in the first iteration,
the length of slice1 is 1. In the second iteration, the length of slice1 is 2, and so on. At line 9,
we are initializing each element at index i with the value i. The length of slice1 is printed in each
iteration at line 11. Finally, the loop at line 14 prints each element of slice1.
*/
