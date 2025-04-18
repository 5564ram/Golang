package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!
	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	// grow the slice:
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	// grow the slice beyond capacity:
	// slice1 = slice1[0:7] // panic: runtime error: slice bounds out of range
}

/*

Slice at 0 is 2
Slice at 1 is 3
Slice at 2 is 4
The length of arr1 is 6
The length of slice1 is 3
The capacity of slice1 is 4
Slice at 0 is 2
Slice at 1 is 3
Slice at 2 is 4
Slice at 3 is 5
The length of slice1 is 4
The capacity of slice1 is 4

Let’s study the code above line by line. In main at line 5, we make an array arr of length 6.
Now, before even initializing the values of arr, we made a slice slice1 out of arr at line 6:var
slice1 []int = arr1[2:5]. The slice1 contains three elements from arr (from index 2 to index 4).
Then, we are initializing elements in the array arr with for loop at line 8. The element at index i
is given the value of the iterator i. So arr holds [0,1,2,3,4,5]. The for loop at line 12 is printing
all the elements from the slice1. You’ll note that although we make slice1 when values are not given
to array arr, slice1 still holds the value from 2nd index to 4th index of arr. So slice1 is [2,3,4].

Line 15 is printing the length of arr, which is 6. Line 16 is printing the length of slice1,
which is 3 (as three elements are present), and line 17 is printing the capacity of slice1,
which is 4 (since three elements are present in slice1 and only one element of arr is beyond
the slice1).

At line 19, we are growing the slice1 as:slice1 = slice1[0:4]. This means that the one last
element of arr1 which was not part of slice1 initially, now is a part of slice1. The for loop
at line 20 is printing all the elements from the updated slice1. Line 23 is printing the length
of slice1, which is 4 (since four elements are present), and line 24 is printing the capacity of
slice1, which is 4 (since four elements are present in slice1 and no element of arr1 is beyond
the slice1).
*/
