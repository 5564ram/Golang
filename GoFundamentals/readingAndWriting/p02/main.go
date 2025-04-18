package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")

	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

/*
The variable inputReader is a pointer to a Reader in bufio. This reader is created and
linked to the standard input in the line: inputReader := bufio.NewReader(os.Stdin).
The bufio.NewReader() constructor has the signature: func NewReader(rd io.Reader) *Reader.
It takes as an argument, any object that satisfies the io.Reader interface (any object
that has a suitable Read() method) and returns a new buffered io.Reader that reads from
the given reader, os.Stdin satisfies this requirement.

The reader has a method ReadString(delim byte), which reads the input until delim is found.
The delim is included; what is read is put into a buffer.

ReadString returns the string read and nil for the error. When it reads until the end of
a file, the string read is returned and io.EOF; if delim is not found, an error err != nil
is returned. In our case, the input from the keyboard is read until the ENTER key
(which contains ‘\n’) is typed. Standard output os.Stdout is on the screen. os.Stderr
is where the error-info can be written to, and it is mostly equal to os.Stdout.

The package bufio allows you to read input from the keyboard or any other source with
its Reader type, which is declared at line 8, and created at line 13. The Reader has all
kinds of methods to read in. The simplest is ReadString, which reads into the input as one
big string until a newline character is found, and returns an error err if there was a
problem (line 16). If no error is found, the input is printed at line 18.

In normal Go-code, the var-declarations are omitted and the variables are declared
with :=, like:

inputReader := bufio.NewReader(os.Stdin)
input, err := inputReader.ReadString('\n')
*/
