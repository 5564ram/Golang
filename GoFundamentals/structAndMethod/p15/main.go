package main

import (
	"fmt"
	"test/person"
)

func main() {
	p := new(person.Person)
	// error: p.firstName undefined
	// (cannot refer to unexported field or method firstName)
	//p.firstName = "Eric"
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName()) // Output: Eric
}

/*
See the person.go file. We have a struct of type Person with two string fields firstName and
lastName at line 3.

Look at the header of the FirstName() method at line 8: func (p *Person) FirstName() string.
From the name, it’s obvious that it is a getter-method. Only the pointer to the object of the
Person type can call it. This method returns the firstName(internal field) of p.

Now, look at the header of the SetFirstName() method at line 12: func (p *Person)
SetFirstName(newName string). From the name, it’s obvious that it is a setter-method.
Only the pointer to the object of the Person type can call it. This method sets the
value of firstName(internal field) of p equal to newName(parameter). Keep in mind that
the type Person can be exported anywhere, but its field can’t be imported.

Let’s look at the main.go file to grasp this concept. At line 4, we import the package person.
In main, at line 8, we make a Person type object p using new() by accessing Person as: person.Person.
We can easily import the type Person with a selector statement, but we can’t export its field.
See the commented line 11, where firstName(internal field of p) is directly accessed.
It will give an error. That’s why we made setter and getter methods.

At line 12, we are using the setter to set the name Eric to the firstName of p. In the next line,
we are using the getter method to get the firstName of p and then printing the returned value.
*/
