package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int // Pointer variable
	intP = &i1    // Storing address of i1 in pointer variable
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}

/*
As you can see, we declared a variable i1 at line 5 and set the value of 5 to it.
In the next line, we print its value along with its memory address with & operator as &i1.
At line 7, we declared a pointer variable intP that is meant to have an address value.
At line 8, we give intP the address of i1 using &i1. Finally, at line 9, we print the address
using intP and value at that address by dereferencing it as *intP, which is 5.
*/
