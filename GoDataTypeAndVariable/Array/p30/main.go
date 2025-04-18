package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)                    // output: [1 2 3 0 0 0 0 0 0 0]
	fmt.Printf("Copied %d elements\n", n) // n == 3
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3) // output: [1 2 3 4 5 6]
}

/*
[1 2 3 0 0 0 0 0 0 0]
Copied 3 elements
[1 2 3 4 5 6]

In the code above, we make two slices: sl_from at line 5 and sl_to at line 6.
The length of sl_from is 3 and the length of sl_to is 10 (all zeros). At line 7,
we copy the contents of sl_from to sl_to and store total number of elements copied in n.
Printing sl_to at line 8 gives [1 2 3 0 0 0 0 0 0 0]. The variable n here is 3 because only
three elements (1,2,3) were copied from sl_from to sl_to. At line 10, we make another slice
called sl3 and initialize it to {1,2,3}. Line 11 shows how to use the append function.
The elements 4, 5 and 6 are further appended at the end of slice sl3. Now, the modified
value of sl3 is {1,2,3,4,5,6}.
*/
