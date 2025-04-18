package main

import "fmt"

func main() {
	s := "good bye"
	var p *string = &s
	*p = "ciao" // changing the value at &s

	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string
}

/*

Here is the pointer p: 0xc82000a270
Here is the string *p: ciao
Here is the string s: ciao

As you can see, we declare a pointer variable p initially and set the address to it of a string variable s, at line 6.
At line 7, we are dereferencing the p variable and changing value to ciao. It means the address that p holds which was &s,
does not have an initial value of s anymore but a new value that is ciao. Then from line 9 to line 11, we are printing to
see the results, which show that dereferencing the p variable actually changes the s value from good bye to ciao.
*/
