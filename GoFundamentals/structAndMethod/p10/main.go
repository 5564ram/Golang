package main

import "fmt"

type S struct {
	a int
}

type SType S        // New type
type SAlias = S     // Alias
type IntType int    // New type
type IntAlias = int // Alias

func (recv S) print() { // function for type defined on type S
	fmt.Printf("%t: %[1]v\n", recv)
}

func (recv SType) print() { // function for type defined on the basis of S
	fmt.Printf("%t: %[1]v\n", recv)
}

// func (recv SAlias) print() { // <-- error: S.print redeclared in this block previous declaration at ./struct_method.go:15:6
// fmt.Printf("%t: %[1]v\n", recv)
// }

func (recv IntType) print() { // function for type defined on type on the basis of int
	fmt.Printf("%t: %[1]v\n", recv)
}

// func (recv IntAlias) print() { // <-- error: cannot define new methods on non-local type int
// fmt.Printf("%t: %[1]v\n", recv)
// }

func main() {
	a := S{10}
	a.print() // calling function from line 13
	b := SType{20}
	b.print() // calling function from line 16
	c := SAlias{30}
	c.print() // calling function from line 13
	d := IntType(40)
	d.print() // calling function from line 24
	// e := IntAlias(50) <-- error: e.print undefined (type int has no field or method print)
	// e.print()
}

/*
In the above code, at line 4, we declare a struct of type S with one int field a.
Then, at line 8, we declare a new type similar to S with the name SType. In the next line,
we alias the type s as SAlias. Similarly, at line 10, we declare a similar type to int as IntType.
In the next line, we alias the type int as IntAlias.

Let’s study the methods involved in our program:

Look at the header of print() method at line 13: func (recv S) print(). The part recv S means that
the object of type S can call this method. This method is printing the value assigned to the field
of the calling object.

Look at the header of the print() method at line 16: func (recv SType) print(). We redefine the print()
method for SType, as this type is defined based on S. So, redefining the method is allowed. The part
recv SType means that the object of type SType can call this method. This method is also printing the
value assigned to the field of the calling object.

See the commented line 20. It is the header for the print() method but also for the object of type SAlias.
This will give an error because for alias, the same method of the base class can’t be redefined.

Look at the header of print() method at line 24: func (recv IntType) print() . The part recv IntType
means that the object of type IntType can call this method. This method is printing the value assigned
to the field of the calling object.

See the commented line 28. It is the function header for the print() method but also for the object
of type IntAlias. This will give an error because for alias, the same method of the base class can’t
be redefined.

Let’s see main now. At line 33, we make an S type object a using struct-literal and assign its
field with a value of 10. In the next line, we call print() method on a, due to which method at
line 13 will get control.

Similarly, at line 35, we made an SType type object b using struct-literal and assign its field
with a value of 20. In the next line, we call print() method on b. Due to which, method at line
16 will gain control.

At line 37, we make an SAlias type object c using struct-literal, and assign its field with a
value of 30. In the next line, we call print() method on c. Due to which, the method at line 13
will gain control (as SAlias is the alias for type S).

At line 39, we make an IntType type object d using struct-literal, and assign its field with a
value of 40. In the next line, we call print() method on d. Due to which, the method at line 24
will gain control.

See the commented line 41, where we made an IntAlias type object e using struct-literal and
assign its field with a value of 50. In the next line, we are calling print() method. It will
generate an error because IntType is an alias for int, where int type doesn’t have any method
called print(). So, it can’t find any method print().
*/
