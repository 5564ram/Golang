package main

import (
	"fmt"
	"strings"
)

func main() {
	var orig string = "Hey, how are you George?"
	var lower string
	var upper string
	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig) // changing to lower case
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig) // changing to upper case
	fmt.Printf("The uppercase string is: %s\n", upper)
}

/*
As you can see in the above code, we declare a string orig and initialize it with Hey, how are you George? , at line 8.
At line 12, we change the case of orig to lowercase, using the function ToLower, and at line 14, we change the case of
orig to uppercase, using the function ToUpper. We print the new strings at line 13 and line 15 to verify the result.

TrimSpace can be used to remove all leading and trailing whitespaces as:

strings.TrimSpace(s)
If you want to trim a specific string str from a string s, use:

strings.Trim(s, str)
For example:

strings.Trim(s, "\r\n")
The above statement will remove all leading and trailing \r and \n from the string s.
The 2nd string-parameter can contain any characters, which are all removed from the left and right-side of s.

If you want to remove only the leading or only trailing characters or strings, use TrimLeft or TrimRight independently.

On whitespaces
The strings.Fields(s) splits the string s around each instance of one or more consecutive white space characters,
and returns a slice of substrings []string of s or an empty list, if s contains only white space.

On a separator
The strings.Split(s, sep) works the same as Fields, but splits around sep.
The sep can be a separator character (:,;,,,-,…) or any separator string sep.

Joining over a slice
The strings.Join(sl []string, sep string) results in a string containing all the elements of the slice sl, separated by sep:

strings.Join([]string{"foo", "bar", "baz"}, ", ")
strings.ToLower(str string) string
strings.ToUpper(str string) string
strings.TrimSpace(str string) string
strings.Trim(str string, cutset string) string
strings.TrimLeft(str string, cutset string) string
strings.TrimRight(str string, cutset string) string
strings.Fields(str string) []string
strings.Split(str string, sep string) []string
strings.Join(sl []string, sep string) string
*/
