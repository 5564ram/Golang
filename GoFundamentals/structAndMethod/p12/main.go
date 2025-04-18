package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }

func main() {
	var b1 B // b1 is a value
	b1.change()
	fmt.Println(b1.write())
	b2 := new(B) // b2 is a pointer
	b2.change()
	fmt.Println(b2.write())
}

/*
In the above code, at line 6, we make a struct of type B with one integer field thing.
Look at the header of change() method at line 10: func (b *B) change(). The part (b *B)
shows that only the pointer to B type object can call this method. This method is changing
its internal field thing by assigning the value of 1 to it. Now, look at the header of the
write() method at line 12: func (b B) write() string. The part (b B) shows that only the B
type object can call this method. This method is returning its internal field thing after
converting it to type string.

Note: The function Sprint() from fmt package returns the string without printing them on the console.

Now, look at main. We make a variable b1 of type B using var keyword at line 15. In the next line,
we call change() method on b1. After returning from this method, thing of b1 will hold value 1.
At line 17, we are printing the result from write() method called on b1. It will print thing from b1,
which will output 1 on the screen. We make a variable b2 of type B using the new() function at line 18.
In the next line, we call the change() method on b2. After returning from this method, thing of b2 will
hold a value of 1. At line 20, we are printing the result from the write() method called on b2, which
will print thing from b2 that will output 1 on the screen.

Notice, in main(), that Go does plumbing work for us; we do not have to figure out whether to call
the methods on a pointer or not, Go does that for us. The variable b1 is a value, and b2 is a pointer.
However, the method calls work just fine. Try to make write() change its receiver value b. You will see
that it compiles fine, but the original b is not changed. We see that a method does not require a pointer
as a receiver as in the following example, where we only need the values of Point3 to compute something:

type Point3 struct { x, y, z float }

// A method on Point3:
func (p Point3) Abs() float {
  return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}
However, this is a bit expensive because Point3 will always be passed to the method by value and copied,
but it is valid in Go. In this case, the method can also be invoked on a pointer to the type
(there is automatic dereferencing). Suppose p3 is defined as a pointer:

p3 := &Point3{3, 4, 5}
Then, you can write p3.Abs() instead of (*p3).Abs().
*/
