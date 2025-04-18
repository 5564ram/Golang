package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{} // empty interface

func main() {
	var val Any
	val = 5 // assigning integer to empty interface
	fmt.Printf("val has the value: %v\n", val)
	val = str // assigning string to empty interface
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1 // assigning *Person type variable to empty interface
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) { // cases defined on type of val
	case int: // if val is int
		fmt.Printf("Type int %T\n", t)
	case string: // if val is string
		fmt.Printf("Type string %T\n", t)
	case bool: // if val is bool
		fmt.Printf("Type boolean %T\n", t)
	case *Person: // if val is *Person
		fmt.Printf("Type pointer to Person %T\n", *t)
	default: // None of the above types
		fmt.Printf("Unexpected type %T", t)
	}
}

/*
val has the value: 5
val has the value: ABC
val has the value: &{Rob Pike 55}
Type pointer to Person main.Person
*/

/*
In the above code, at line 4 and line 5, we make two global variables i and str, respectively,
and set their values. Then, at line 7, we define a Person type struct with two fields name (a string)
and age (an integer). Next, at line 12, we make an empty interface Any.

Now, look at main. We make Any type variable val at line 15. In the next line, we set val to 5.
At line 18, val gets the value of str, so in the next line, ABC (value of str) will be printed.

At line 20, we make pers1 a Person pointer-type variable. In the next two lines, we give a value
to the internal fields ( Rob Pike as name and 55 as age) of pers1. At line 23, val1 is given the
value of pers1. In the next line, val is printed, which prints the pers1 struct as &{Rob Pike 55}.

Now, we have switch cases ahead. Cases are to be judged on type of val.

-> The first case will be true if val is of int type.
-> The second case will be true if val is of string type.
-> The third case will be true if val is of bool type.
-> The fourth case will be true if val is of *Person type.
Now, the type will be printed accordingly when one case is found true. In this case, the last updated
value of val was pers1 of type *Person. So, the third case will become true, and line 33 will be executed.

What if, no case is found true? In this case, we have a default case that prints the message of the
unexpected type.

Each interface{} variable takes up 2 words in memory: one word for the type of what is contained,
and the other word for either the contained data or a pointer to it.
*/
