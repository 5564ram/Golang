package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? \nDoes the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("\n%t\n\n", strings.HasPrefix(str, "Th")) // Finding prefix

	fmt.Printf("Does the string \"%s\" have suffix %s? ", str, "ting")
	fmt.Printf("\n%t\n\n", strings.HasSuffix(str, "ting")) // Finding suffix
}

/*
As you can see in the above code, we declare a string str and initialize it with This is an example
of a string at line 9. At line 11, we used the function HasPrefix to find prefix Th in string str.
The function returns true because str does start with Th. Similarly, at line 14, we used the function
HasSuffix to find the suffix ting in string str. The function returns false because str does not end
with ting. This also illustrates the use of the escape character \ to output a literal " with \" , and
the use of 2 substitutions in a format-string.

HasPrefix(s, prefix string) bool
HasSuffix(s, suffix string) bool
*/
