package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1
	// shorter, without separate declaration:
	// areaIntf := Shaper(sq1)
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}

/*
In the program above, we define an interface Shaper at line 4, with one method Area() returning
float32 value and a struct of type Square at line 8, with one field side of type float32. At line 12,
we define a method Area() that can be called by a pointer to the Square type object. This method
returns the area of a square sq. The struct Square implements the interface Shaper. Now, the interface
variable contains a reference to the Square variable, and through it, we can call the method Area() on Square.

Look at the main. We make a Square type variable sq1 at line 17 and give a value of 5 to its field side
(at line 18). As discussed above, you could call the method immediately on the Square value sq1.Area(),
but the novel thing is that we can call it on the interface value (see line 23), thereby generalizing the call.

The interface variable both contains the value of the receiver value and a pointer to the appropriate
method in a method table.

This is Go’s version of polymorphism, a well-known concept in OO software. The right method is chosen
according to the current type or put otherwise. A type seems to exhibit different behaviors when linked
to different values. If Square didn’t have an implementation of Area(), we would receive the clear
compiler error: cannot use sq1 (type *Square) as type Shaper in assignment: *Square does not implement
Shaper (missing Area method).

The same error would occur if Shaper had another method Perimeter(), and Square would not have an
implementation for that. We expand the example with a type Rectangle which also implements Shaper.
Now, we can make an array with elements of type Shaper and show polymorphism in action by using a
for range on it and calling Area() on each item:
*/
