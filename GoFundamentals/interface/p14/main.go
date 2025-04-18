package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// setting a value:
	// Error: will panic: reflect.Value.SetFloat using unaddressable value
	// v.SetFloat(3.1415)
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())

	fmt.Println(v)
}

/*
settability of v: false
type of v: *float64
settability of v: false
The Elem of v is:  3.4
settability of v: true
3.1415
3.1415
*/

/*
Continuing with the previous example, suppose we want to modify the value of x to be 3.1415.
Value has a number of Set methods to do this, but here we must be careful:
v.SetFloat(3.1415)
It produces an error: will panic: reflect.Value.SetFloat using unaddressable value. Why is
this? The problem is that v is not settable (not that the value is not addressable).
Settability is a property of a reflection Value, and not all reflection Values have it.
It can be tested with the CanSet() method. In our case, we see that this is false, settability
of v: false. When v was created with v := reflect.ValueOf(x), a copy of x was passed to the
function, so it is logical that you can’t change the original x through v.
In order to change x through v, we need to pass the address of x: v = reflect.ValueOf(&x).
Through Type(), we see that v is now of the type *float64 and still not settable. To make it
settable, we need to let the Elem() function work on it, which indirects through the pointer
v = v.Elem(). Now, v.CanSet() gives true and v.SetFloat(3.1415) works!
*/

/*
In the above code, at line 4, we import a package reflect. At line 8, we declare a float64 type
variable x and assign 3.4 value to it. At line 9, we are reading the value of x in the new variable
v through ValueOf() function of the reflect package.
Now, before setting a value to x it’s better to check it’s settable or not to avoid errors. Look at
line 13. We are calling canSet() method on v, and printing the result. It will print false, as v isn’t
settable because it’s unaddressable. Look at line 14, we are reading the address of x(pointer) in new
variable v through ValueOf() function of the reflect package. In the next line, printing the type of
v gives *float64. Again, before setting a value to x, it’s better to check if it’s settable or not to
avoid errors. Look at line 16. We are calling canSet() method on v, and printing the result. It will
print false, as v isn’t settable.

Now, we try Elem() method on v at line 17. The line 18 prints the elem of v, which is its value 3.4.
In the next line, we are again checking its settability. This time it works. We set the value of v to a
floating value 3.1415 at line 20. At line 21, we recover the interface value of v, and print v at the line
23, which is 3.1415 now.

Reflection Values need the address of something in order to modify what they represent.
*/
