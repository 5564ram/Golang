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

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter:
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}

/*
Looping through shapes for area ...
Shape details:  {5 3}
Area of this shape is:  15
Shape details:  &{5}
Area of this shape is:  25
*/
/*
In the above program, we define an interface Shaper at line 4 with one method Area(), returning a
float32 value and a struct of type Square at line 8 with one field side of type float32. At line 12,
we define a method Area() that can be called by a pointer to the Square type object. Again at line 16,
we make a struct of type Rectangle at line 16 with one field side of type float32. At line 20, we define
a method Area() that can be called by the Rectangle type object.

Now, look at main. We make a value type Rectangle object and a pointer type Square object. As the interface
variable contains the reference to the Square and Rectangle type object, we create a Shaper type array shapes
and add r (a Rectangle object) and q (a Square object) in it.

Our goal is to find the area of all the shapes present in shapes. We use a for loop at line 31. For each
iteration, we print the details of a shape in shapes (see line 32) and print the area (see line 33).
In this case, our first shape is the rectangle r. So, line 32 will output {5,3} on the screen. At line 33,
Area(), having the object of Rectangle as a receiver (see line 20), will be called, and 15 will be returned
from the method and printed on the screen by line 33. In the next iteration, our shape is the square q. So,
line 32 will output &{5} on the screen, as q is the pointer variable. At line 33, Area(), having a pointer
to the object of Square as a receiver (see line 12), will be called, and 25 will be returned from the method
and printed on screen by line 33.
*/
