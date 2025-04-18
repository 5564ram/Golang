package main

import "fmt"

func main() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
}

/*
We see that with range at line 6 the returned character c from s is in Unicode, so it is a rune.
Unicode-characters take 2 bytes; some characters can even take 3 or 4 bytes. If erroneous UTF-8 is encountered,
the character is set to U+FFFD and the index advances by one byte.

In the same way as above, the conversion:

c:=[]byte(s)
is allowed. Then, each byte contains a Unicode code point, meaning that every character from the string corresponds to one byte.
Similarly, the conversion to runes can be done with:

r:=[]rune(s)
The number of characters in a string s is given by len([]byte(s)). A string may also be appended to a byte slice, like this:

var b []byte
var s string
b = append(b, s...)

Making a substring of a string
The following line:

substr := str[start:end]
takes the substring from str from the byte at index start to the byte at index end– 1.
Also, str[start:] is the substring starting from index start to len(str) – 1, and str[:end] is the substring starting from index 0 to end – 1.

------------- Changing a character in a string# -------------
Strings are immutable. This means when str denotes a string then str[index] cannot be the left side of an assignment:

str[i] = 'D'
Where i is a valid index, it gives the error cannot assign to str[i].

To change a character, you first have to convert the string to an array of bytes. Then, an array-item of
a certain index can be changed, and the array must be converted back to a new string. For example, change “hello” to “cello”:

s := "hello"
c := []byte(s) // converting string s to array of bytes `c`.
c[0] = 'c'     // modifying the c
s2 := string(c) // Converting c to string, s2 == "cello"
It is clear that string-extractions are very easy with the slice-syntax. But changing a character in a string is not as easy.
*/
