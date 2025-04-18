package main

import "fmt"

type struct1 struct { // struct definition
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1) // making a struct1 type variable
	// Filling fields of the struct of struct1 type
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"
	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)
}

/*
In the above program, look at line 4. We are declaring a struct of type struct1. It has three different fields.
First, we have an integer i1(see line 5), then we have a float32 type variable f1(see line 6) and lastly, there
is a string variable str (see line 7). This means, if we made a variable of type struct1, it would hold these three different
fields to which we can assign values by our own choice. Now, look at the main function. At line 11, we make a struct1 type
struct ms using the new function as: ms := new(struct1). In the next three lines, we are assigning values to the fields of ms.
At line 13, we assign ms.i1 the value of 10. At line 14, we assign ms.f1 the value of 15.5. At line 15, we assign ms.str the
value of Chris. In the next lines (from line 16 to line 19), we are printing the fields’ values for ms to verify the results.

In the following example, we see that you can also initialize values by preceding them with the field names.
So, new(Type) and &Type{} are equivalent expressions. A typical example of a struct is a time-interval
(with a start and end time expressed here in seconds):

type Interval struct {
  start int
  end int
}
Here, are some initializations:

inter := Interval{0,3} //(A)
inter2 := Interval{end:5, start:1} //(B)
inter3 := Interval{end:5} //(C)

In case A, the values given in the literal must be exactly in the same order as the fields are defined in the struct;
the & is not mandatory. Case B shows another possibility where the field names with a : precede the values; in that case,
the order of the field names can be different from the order of their definition. Also, if fields are omitted like in case C,
zero-values are assigned. The naming of the struct type and its fields adheres to the Visibility-rule. An exported
struct type may have a mix of fields where some are exported and others not.
*/
