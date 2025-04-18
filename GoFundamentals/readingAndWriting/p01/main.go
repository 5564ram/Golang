package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}

/*
The Scan functions require us to first define the variables in which the input is to be stored.
So from line 5 to line 7, we declare a total of 5 variables:

At line 5, we have three string variables: firstname, lastname, and s.
At line 6, we have one integer variable: i.
At line 7, we have one float32 type variable: f.
At line 14, we use fmt.Scanln to read in a line from the keyboard until the user types ENTER.
We asked the user to type in their full name so in the format firstname SPACE lastname.
These strings will be stored in the same order in the parameters to Scanln:
fmt.Scanln(&firstName, &lastName). Scanln needs references to the storage variables, that’s why
it takes the parameters &firstName and &lastName. To verify that, we print out the variables at
line 16.

Go has also the Scanf, Fscanf, and Sscanf functions that parse the arguments according to a
format string, analogous to that of Printf. Scanf reads input from the keyboard, but Sscanf
reads from another string.

At line 17, we see that Sscanf takes the following parameters, which were also defined in the
variables section:

input: the string to be read
format: the format to read and parse input
This contains (a list of) the storage variables &f, &i, and &s for storing the input. So,
the input e.g., 56.12 / 5212 / Go is parsed according to format %f / %d / %s and stored in
the variables &f, &i, and &s, respectively.

Scanln scans the text read from standard input, storing successive space-separated values
into successive arguments; it stops scanning at a newline. Scanf does the same, but it uses
its first parameter as the format string to read the values in the variables. Sscan and
friends work the same way but read from an input string instead of reading from standard
input. You can check on the number of items read in and the error when something goes wrong.
*/
