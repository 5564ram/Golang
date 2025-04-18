package main

import "fmt"

func main() {
	s2 := []int{1, 2, 3, 4, 5}
	s3 := s2[1:] // Creates a new slice starting from index 1

	s3[0] = 99 // Modifies the underlying array

	fmt.Println("s2:", s2) // Output: s2: [1 99 3 4 5]
	fmt.Println("s3:", s3) // Output: s3: [99 3 4 5]
}

/*
Slices in Go are reference types.

s3 does not create a new independent copy; it shares memory with s2.

Modifying elements in s3 affects s2 because they point to the same underlying array.

If you want an independent copy, you need to explicitly copy the elements:
s3 := append([]int{}, s2[1:]...) // Creates a new slice with copied elements
*/
