package main

import (
	"fmt"
	"strings"
)

func main() {
	var origS string = "Hi there! "
	// var newS string
	newS := strings.Repeat(origS, 3) // Repeating origS 3 times
	fmt.Printf("The new repeated string is: %s\n", newS)
}

/*
As you can see in the above code, we declare a string origS and initialize it with Hi there! at line 8.
At line 10, we repeat origS 3 times using the Repeat function and store the new string in newS at line 10.
Then, the result is printed at line 11.
strings.Repeat(str, count int) string
*/
