package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of the first instance of\"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc")) // Finding first occurence
	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi")) // Finding first occurence
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi")) // Finding last occurence
	fmt.Printf("The position of the first instance of\"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger")) // Finding first occurence
}

/*
As you can see in the above code, we declare a string str and initialize it with Hi, I’m Marc, Hi.
at line 8. At line 10, we find the index of the first occurrence of Marc in str using the Index function
that gives 8 as a result. Similarly, at line 12, we find the index of the first occurrence of Hi in str
using the Index function, which is 0, and its last occurrence by using LastIndex at line 14, which is 14.
At line 16, we find the first occurrence of Burger in str which gives -1 because it doesn’t exist.

strings.Replace(str, old, new string, n int)
We can replace the first n occurrences of old in str by new.
A copy of str is returned, and if n is -1, all occurrences are replaced.

strings.Index(str, substr string) int
strings.LastIndex(str, substr string) int
strings.IndexRune(str string, r rune) int
strings.Replace(str, old, new string, n int)
*/
