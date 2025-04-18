package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
}

/*
The output of this example can be different every time because the value &i1 can be different each time we execute this statement.

This address can be stored in a special data type called a pointer. In the above case, it is a pointer to an int. So i1 is denoted by *int.
If we call that pointer intP, we can declare it as:

var intP *int
Then, the following statement is true:

intP = &i1
So, intP stores the memory address of i1 which means it points to the location of i1. In other words, it references the variable i1.

A pointer variable contains the memory address of another value which means it points to that value in memory.
The size of a pointer is 4 bytes on 32-bit machines, and 8 bytes on 64-bit machines, regardless of the size of
the value they point to. Of course, pointers can be declared to a value of any type, be it primitive or structured.
The * is placed before the type of the value (prefixing), so the * is a type modifier here.

Using a pointer to refer to a value is called indirection. A newly declared pointer which has not been assigned to
a variable has the nil value. A pointer variable is often abbreviated as ptr.

In an expression like:

var p *type
always leave a space between the name of the pointer and the *. The statement var p*type is syntactically correct, but
in more complex expressions, it can easily be mistaken for a multiplication.

The same symbol * can be placed before a pointer like *intP, and it gives the value which the pointer is pointing to
and it is called the dereference (or contents or indirection) operator. Another way to say this is that the pointer is flattened.
So for any variable var, the following is true:

var == *(&var)
*/
