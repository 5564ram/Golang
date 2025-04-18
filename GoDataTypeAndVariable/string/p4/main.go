package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"
	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H")) // count occurences
	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg")) // count occurences
}

/*
As you can see in the above code, we declare the two strings str and manyG and initialize them with Hello,
how is it going, Hugo? and gggggggggg at line 8 and line 9, respectively. At line 11, we find the count of H
in str using the Count function that gives 2 as a result. Similarly, at line 13, we find the count of gg in
manyG using the Count function, which is 5
strings.Count(str, substr string) int
The function returns the number of non-overlapping instances of substr in str.
*/
