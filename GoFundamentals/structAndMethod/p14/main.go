package main

import "fmt"

type T struct {
	a int
}

func (t T) print(message string) {
	fmt.Println(message, t.a)
}

func (T) hello(message string) {
	fmt.Println("Hello!", message)
}

func callMethod(t T, method func(T, string)) {
	method(t, "A message")
}

func main() {
	t1 := T{10}
	t2 := T{20}
	var f func(T, string) = T.print
	callMethod(t1, f)
	callMethod(t2, f)
	callMethod(t1, T.hello)
}

/*
Methods as values and parameters
Methods are just like functions in that they can be used as values, and passed to other
functions as parameters. This is shown in the following program:
*/

/*
A message 10
A message 20
Hello! A message

In the above program, at line 4, we make a struct of type T with a single integer field a.
Look at the header of the print method at line 8 as: func (t T) print(message string).
The part (t T) means that this method can only be called by an object of type T. This method
is printing message sent as a parameter, and the internal field a of t. Look at the header
of the print method at line 12 as: func (T) hello(message string). The part (T) means that
this method can be directly called with type T as: T.hello("anyString"). This method is
printing Hello!, and then prints the message parameter. Look at the header of callMethod
function at line 16 as: func callMethod(t T, method func(T, string)). This function takes
t as a parameter and a function method. That function method is called with T and a string
“A message” as parameters.

Now, look at main. We made a T type variable t1 using struct-literal giving a value of 10,
at line 21. Similarly, in the next line, we make a T type variable t2 using struct-literal
giving a value of 20. At line 23, we make a variable f equal to the print method of T.

Now, at line 24, we are calling the callMethod function with t1 and f as parameters. Here,
the method in callMethod is equal to the print() method of type T. So, from line 17, print() '
will be called for t1 with message as “A message”. So, A message 10 will be printed on the screen.
Similarly, in the next line, we are calling the callMethod function with t2 and f as parameters.
Here, the method in callMethod is equal to print() method of type T. So, from line 17, print()
will be called for t2 with message as “A message”. So, A message 20 will be printed on the screen.

Now at line 26, we are calling the callMethod function with t1 and T.hello as parameters. Here
the method in callMethod is equal to the hello() method of type T. So, from line 17, hello()
will be called for t1 with the message as A message. So, Hello! A message will be printed on the screen.
*/
