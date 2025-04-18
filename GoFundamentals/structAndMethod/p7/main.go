package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32

	int    // anonymous field
	innerS // anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// with a struct-literal:
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is: ", outer2)
}

/*
outer.b is: 6
outer.c is: 7.500000
outer.int is: 60
outer.in1 is: 5
outer.in2 is: 10
outer2 is:  {6 7.5 60 {5 10}}

In the above code, at line 4, we make a struct of type innerS containing two integer fields in1 and in2. At line 9,
we make another struct of type outerS with two fields b (an integer) and c (a float32 number). That’s not all.
We also have two more fields that are anonymous fields. One is of type int (see line 13), and the other is of type innerS (see line 14).
They are anonymous fields because they have no explicit names.

Now, look at main. We make an outerS variable outer via the new function at line 18. In the next few lines, we are giving the values to
its fields. The fields b and c of outer are given values at line 19 and line 20, respectively. Now, it’s the turn for anonymous fields of outer.

To store data in an anonymous field or get access to the data, we use the name of the data type, e.g. outer.int. A consequence
is that we can only have one anonymous field of each data type in a struct. Look at line 21. We are assigning the value of 60
to the int type anonymous field declared at line 13. But how to assign the value to the second anonymous field innerS declared
at line 14? The answer is simple. We also have a struct of type innerS in our program, which contains two int fields in1 and in2.
So, to assign the value to the second anonymous field innerS, we’ll assign the values to the fields of innerS( see line 22 and line 23).
From line 24 to line 28, we are printing the above fields of outer to which we assigned the values, to verify the results.

This is one of the methods to make and assign values to a struct type variable that contains anonymous fields, i.e., using the
new function and then the selector operator. How about making such a struct type variable using struct-literal? Look at line 30,
we are making an outerS variable outer2, and assigning the same values as the above but in a struct-literal manner as:
outer2 := outerS{6, 7.5, 60, innerS{5, 10}}. The order is followed. The variable b gets 6, and c gets 7.5. The first anonymous
field of outer2, which is an int, gets 60. The anonymous field of innerS gets the value 5 for in1 and 10 for in2. In the last
line, we are printing outer2 to check the values assigned to its fields.

To store data in an anonymous field or get access to the data, we use the name of the data type, e.g., outer.int.
A consequence is that we can only have one anonymous field of each data type in a struct.
*/
