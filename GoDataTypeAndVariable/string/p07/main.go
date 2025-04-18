package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, _ = strconv.Atoi(orig) // converting to number
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an) // converting to string
	fmt.Printf("The new string is: %s\n", newS)
}

/*
As you can see in the above code, we declare a string orig and initialize it with 666 ,
at line 8. At line 11, we calculate the size of ints using strconv.IntSize that appears to be 64.
At line 12 we convert orig to an integer using the function strconv.Atoi(orig) and store it in a new variable an.
We modify an and convert it into a string at line 15 using strconv.Itoa(an) and then store it in a new variable newS.
an and newS are printed at line 13 and line 16 to verify the results.
https://pkg.go.dev/strconv
*/
