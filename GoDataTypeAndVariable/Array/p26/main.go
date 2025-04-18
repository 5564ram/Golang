package main

import "fmt"

func main() {

	// -----new-----
	p := new(int)   // Allocates memory for an int, initialized to zero
	fmt.Println(*p) // Output: 0

	s := new([]int)       // Allocates a pointer to a slice, but the slice itself is nil
	fmt.Println(s == nil) // Output: false (s is a pointer, but its value is nil)

	// -----make-----
	s1 := make([]int, 5) // Creates a slice with length 5
	fmt.Println(s1)      // Output: [0 0 0 0 0]

	m := make(map[string]int) // Creates an empty map
	m["age"] = 25
	fmt.Println(m) // Output: map[age:25]

	c := make(chan int, 2) // Creates a buffered channel
	fmt.Println(c)         // Output: a valid channel reference
}

/*

new()
Purpose: Allocates zeroed memory for a value of a given type and returns a pointer to it.
Used With: Pointers to values (e.g., structs, arrays, primitive types).
Return Type: A pointer to the allocated memory (*T).

make()
Returns an initialized value (T), not a pointer.
Only works with slices, maps, and channels.
Allocates memory and initializes the structure.

Feature				new(T)             vs              						make(T)

Works With  	Any type (structs, arrays, primitives)					Only slices, maps, and channels

Returns     	Pointer to allocated memory (*T)						An initialized instance (T)

Initialization	No (only zeroed memory)									Yes (fully initialized)

Used For		Simple memory allocation								Creating and initializing slices, maps, or channels

*/
