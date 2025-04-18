package main

import (
	"fmt"
	"time"
)

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("The sum is: %d\n", two1.AddThem())                 // calling method
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20)) // calling method
	two2 := &TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem()) // calling method
	fmt.Println("a: ", two2.a, "b: ", two2.b)
	main1()
}

func (tn *TwoInts) AddThem() int { // can be called by pointer to TwoInt type var.
	tn.a = tn.a + 10
	tn.b = tn.b + 10
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int { // can be called by pointer to TwoInt type var.
	return tn.a + tn.b + param
}

/*
func (tn *TwoInts) method(), TwoInts       reason -> Go automatically converts two2 to &two2
func (tn *TwoInts) method(), *TwoInts      reason -> Normal value method call
func (tn TwoInts) method(), TwoInts	       reason -> Normal value method call
func (tn TwoInts) method(), *TwoInts       reason -> Go automatically dereferences *two2 to two2
*/

/*
In the above code, at line 4, we make a struct TwoInts with two integer fields a and b.
Before going in main, let’s see two basic methods: AddThem() and AddtoParam. Look at the
header of method AddThem at line 19: func (tn *TwoInts) AddThem() int. The part (tn *TwoInts)
means that the pointer to the variable of type TwoInts can call this method. It takes nothing
as a parameter and returns an int value. It adds both fields of the variable to which tn is
pointing using the selector operator (see line 20), and return the sum. Similarly, at line 23,
look at the header of the method AddtoParam: func (tn *TwoInts) AddToParam(param int) int.
The part (tn *TwoInts) means that pointer to the variable of type TwoInts can call this method.
It takes an integer param as a parameter, and returns an int value. It adds both fields of the
variable to which tn is pointing using the selector operator and then adds param also (see line 24)
and returns the sum.

Now, look at main. At line 10, we are making a TwoInts type variable two1. In the next two lines,
we are assigning the values to the fields of two1. At line 13, we are calling the method AddThem
on two1 as: two1.AddThem. It will return 22 as a of two1 is 12 (see line 11), and b of two1 is 10
(see line 12). Similarly, in the next line, we are calling the method AddToParamon two1 as:
two1.AddToParam(20). It will return 42 (10+12+20). At line 15, we are making a TwoInts type variable
two2 using struct-literal, giving a of two2 the value of 3 and b of two2 the value of 4. At line 16,
we are calling method AddThemon two2 as: two2.AddThem. It will return 7 (3+4).

A method and the type on which it acts must be defined in the same package; that’s why you cannot
define methods on type int, float, or the like. Trying to define a method on an int type gives the
compiler error: cannot define new methods on non-local type int.

For example, if you want to define the following method on time.Time:

func (t time.Time) first3Chars() string {
return time.LocalTime().String()[0:3]
}

You get the same error for a type defined in another, thus also non-local package.
However, there is a way around this: you can define an alias for that type (int, float, …),
and then define a method for that alias. Or, embed the type as an unknown type in a new struct,
like in the following example. Of course, this method is only valid for the alias type.
*/

type myTime struct {
	time.Time //anonymous field
}

func (t *myTime) first3Chars() string {
	return t.String()[0:3]
}

func main1() {
	m := myTime{time.Now()}
	//calling existing String method on anonymous Time field
	fmt.Println("Full time now:", m.String())
	fmt.Println("First 3 chars:", m.first3Chars()) //calling myTime.first3Chars
}

/*
In the above code, we importe package time at line 4 to use its methods. At line 7,
we make a struct of type myTime and create an anonymous field of type time.Time in it.
Look at the header of method first3Chars() at line 11: func (t myTime) first3Chars() string .
The part (t myTime) means that the variable of type myTime can call this method. It takes nothing
as a parameter and returns a string value. It converts the time stored in the field of t to string
and returns the first three characters (see line 12).

Now, look at main. At line 16, we are creating a variable of type myTime m using struct-literal,
and setting its anonymous field of type time.Time to the present time. At line 18, we are printing
the string value for m. It means that the present time that was stored at line 16 will be printed
in the form of a string. In the next line, we call the method first3Chars on m as: m.first3Chars(),
which will print only the first three characters of the present time from line 16 after returning
from the method.

*/
