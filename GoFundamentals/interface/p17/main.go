package main

import "fmt"

type Shaper interface {
	Area() float32
}

type TopologicalGenus interface {
	Rank() int
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Rank() int {
	return 1
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (r Rectangle) Rank() int {
	return 2
}

func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
	topgen := []TopologicalGenus{r, q}
	fmt.Println("Looping through topgen for rank ...")
	for n := range topgen {
		fmt.Println("Shape details: ", topgen[n])
		fmt.Println("Topological Genus of this shape is: ", topgen[n].Rank())
	}
}

/*
For the program above we found out that, we need a new interface TopologicalGenus, defined at line 9,
which gives the rank of a shape (here, simply implemented as returning an int).

At line 40, we create an array shaper of type Shaper and populate it with r and q. Then, from line 42
to line 44, we loop through this array and call the Area() method on both.

In the exactly same way, we can now use our new TopologicalGenus type: at line 46, we create an array
topgen of type TopologicalGenus, and populate it with r and q. Then, we loop through this array and
call the Rank() method on both.

Notice that in order to introduce this new interface, we didn’t need to change the Square and Rectangle
types themselves. We only needed to make them implement the method Rank() of TopologicalGenus. You may
have noticed that for Square the method Rank() is implemented which returns 1 (from line 21 to line23).
Similarly for Rectangle, the method Rank() is implemented which returns 2 (from line 33 to line35).

So you don’t have to work all your interfaces out ahead of time; the whole design can evolve without
invalidating early decisions. If a type must implement a new interface, the type itself doesn’t have to
be changed; you must only make the new method(s) on the type.
*/
