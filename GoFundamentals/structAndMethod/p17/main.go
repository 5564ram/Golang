package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point // anonymous field of Point type
	name  string
}

func main() {
	n := &NamedPoint{Point{3, 4}, "Pythagoras"} // making pointer type variable
	fmt.Println(n.Abs())                        // prints 5
}

/*
In the above code, at line 7, we make a struct of type Point with two fields x and y of type float64.
We make another struct at line 15, of type NamedPoint with two fields in it. The first is an anonymous
field of type Point and the second is name, a variable of type string. Look at the header of the method Abs()
at line 11: func (p *Point) Abs() float64 . It shows that this method can only be called by the pointer to
the variable of type Point, and it returns a value of type float64. Following is the formula for calculating
the absolute value of a point:

value= sqrt(x^2+y^2)
Line 12 of the code implements the above formula. We get the values of x and y through the selector operator
applied on Point p. To calculate the square root, we use the package of math which is imported at line 4.

The method calculates the absolute value for a point and returns it. Now, look at main. At line 21, we make
a pointer variable n of type NamedPoint using the struct-literal: &NamedPoint{Point{3, 4}, "Pythagoras"}.
The part Point{3,4} is used to assign value to the anonymous Point variable(x gets 3 and y gets 4), and
Pythagoras is assigned to the name string. The symbol & makes it a pointer variable. At line 22, we call
the method Abs on variable n and print the return value, which is 5

Embedding injects fields and methods of an existing type into another type. Of course, a type can have
methods that act only on variables of that type, not on variables of the embedded parent type.
*/
