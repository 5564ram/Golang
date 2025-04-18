package main

import "fmt"

type Camera struct{}

func (c *Camera) TakeAPicture() string { // method of Camera
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string { // method of Phone
	return "Ring Ring"
}

// multiple inheritance
type SmartPhone struct { // can use methods of both structs
	Camera
	Phone
}

func main() {
	cp := new(SmartPhone)
	fmt.Println("Our new SmartPhone exhibits multiple behaviors ...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
}

/*
Multiple inheritance is the ability for a type to obtain the behaviors of more than one parent class.
In classical OO languages, it is usually not implemented (exceptions are C++ and Python), because,
in class-based hierarchies, it introduces additional complexities for the compiler. But in Go, multiple
inheritance can be implemented simply by embedding all the necessary ‘parent’ types in the type under
construction.
*/

/*
In the above code, at line 4, we make a struct of type Camera with no fields. At line 6, there is a method
that can be called only by a pointer that points to a Camera object. The method TakeAPicture returns a string
Click. At line 10, we make a struct of type Phone with no fields. At line 12, there is a method that can be
called only by a pointer pointing to a Phone object. The method Call just returns a string Ring Ring.

Now, we have a struct SmartPhone at line 17. It can implement the functionalities of both Camera and Phone.
It means we need multiple inheritance now. So, we made two anonymous fields in SmartPhone: Camera and Phone.
Now, look at main. We make a SmartPhone type pointer variable cp using the new() function. Now, this cp object
can call methods of Camera and Phone as well. See line 25. We are calling the method TakeAPicture, which can
be called by the Camera object. Since SmartPhone inherits Camera, cp.TakeAPicture() works fine. It will print
Click on the screen. In the next line, we are calling the method Call(), which can be called by the Phone
object. Since SmartPhone inherits Phone, cp.Call() works fine. So it will print Ring Ring on the screen.
*/
