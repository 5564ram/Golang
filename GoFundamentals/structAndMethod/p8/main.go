package main

import "fmt"

func main() {

	var person struct {
		name, surname string
	}

	person.name, person.surname = "Barack", "Obama"
	anotherPerson := struct { // anonymous struct
		name, surname string
	}{"Barack", "Obama"}
	fmt.Println(person, anotherPerson)
}

/*
In the above code, at line 6, we make a struct type person with two string fields name and surname.
At line 10, we set the fields of the person struct. The field name gets Barack and surname gets Obama.
At line 11, we make an anonymous struct anotherPerson. We do not define any struct type alias but create
a struct from the inline struct type and define the initial values in the same syntax. We assign the value
to name and surname as Barack and Obama. At line 14, we are printing person and anotherPerson. The output
shows that they both print the same result.

What are the rules when there are two fields with the same name (possibly a type-derived name) in the outer struct and an inner struct?

1. An outer name hides an inner name. This provides a way to override a field or method.
2. If the same name appears twice at the same level, it is an error if the program uses the name. If it’s not used, it doesn’t matter.
There are no rules to resolve the ambiguity; it must be fixed. For example:

type A struct { a int }
type B struct { a, b int }
type C struct { A; B }
var c C
According to rule (2), when we use c.a, it is an error because there is ambiguity on what is meant: c.A.a or c.B.a.
The compiler error is: ambiguous DOT reference c.a. We have to disambiguate with either c.A.a or c.B.a. Look at another example:

type D struct { B; b float32 }
var d D

According to rule (1), using d.b is ok. It is the float32, not the b from B. If we want the inner b, we can get at it by d.B.b.
*/
