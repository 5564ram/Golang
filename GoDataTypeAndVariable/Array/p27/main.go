package main

import (
	"bytes"
	"fmt"
)

func main() {
	// New Buffer.
	var b bytes.Buffer
	// Write strings to the Buffer.

	b.WriteString("ABC")
	b.WriteString("DEF")
	// Use Fprintf with Buffer.
	fmt.Fprintf(&b, " A number: %d, a string: %v\n", 10, "bird")
	b.WriteString("[DONE]")
	// Convert to a string and print it.
	fmt.Println(b.String())
}

/*
in the code above, in main at line 9, we make a buffer as: var b bytes.Buffer.
Now, we start writing to the buffer b. Look at line 12 and line 13, where we write two strings
ABC and DEF to the buffer b. At line 15, we are printing the value placed at address &b,
which will print ABCDEF although written to b initially. Then at line 16, we write [DONE] to
the buffer b and print it by b.String(), which returns string written to the buffer at line 18.

A Buffer can be created as a value, like:
var buffer bytes.Buffer

or as a pointer with new as:

var r *bytes.Buffer = new(bytes.Buffer)
or created with the function:

func NewBuffer(buf []byte) *Buffer
*/
