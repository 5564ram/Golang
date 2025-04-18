package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	// Is Square the type of areaIntf ?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}

/*
The type of areaIntf is: *main.Square
areaIntf does not contain a variable of type Circle
*/

/*
In the above code, at line 7, we make a struct of type Square with one field side of type float32.
Similarly, at line 11, we make a struct of type Circle with one field radius of type float32.
At line 15, we make an interface Shaper with a single method Area(). At line 35, we define a method Area()
that can be called by a pointer to the Square type object. Again, at line 39, we define a method Area()
that can be called by a pointer to the Circle type object.

Now, look at main. We create a Shaper variable areaIntf at line 20. In the next line, we make a Square
variable sq1 and assign it to areaIntf. At line 25, we check whether areaIntf contains a reference to
the type Square using the if condition. If ok is true, control will transfer to line 26, and its type
will be printed. Otherwise, control will transfer to line 28. At line 28, we are checking whether areaIntf
contains a reference to the type Circle using the if condition.

If ok is true, control will transfer to line 29, and its type will be printed. Otherwise, control will
transfer to line 31, and a message will be printed on the screen that such type doesn’t exist.

Remark: If we omit the * in areaIntf.(*Square), we get the compiler-error: impossible type assertion:
areaIntf (type Shaper) cannot have dynamic type Square (missing Area method).
*/
