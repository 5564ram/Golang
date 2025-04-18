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

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)

	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)

	default:
		fmt.Printf("Unexpected type %T", t)
	}
}

/*
Type Square *main.Square with value &{5}
*/
/*
The variable t receives both value and type from areaIntf. All of the listed types have to implement
the interface (Shaper in this case); if the current type is none of the case-types, the default clause
is executed. Fallthrough is not permitted. With a type-switch, a runtime type analysis can be done.
Of course, all the built-in types as int, bool, and string can also be tested in a type switch.

Never use element.(type) outside of a switch statement.
*/
