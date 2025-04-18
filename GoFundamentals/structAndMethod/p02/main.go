package main

import (
	"fmt"
	"strings"
)

type Person struct { // struct definition
	firstName string
	lastName  string
}

func upPerson(p *Person) { // function using struct as a parameter
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1- struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2 - struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" // this is also valid
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3 - struct as a literal:
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}

/*
In the above program, look at line 7. We are declaring a struct of type Person. It has two different fields:
a string called firstName(see line 8) and a string called lastName(see line 9). Look at the header of the function
upPerson at line 12: func upPerson (p *Person). Here, a pointer to the Person type variable is passed as a parameter,
which means if the values of the fields are changed inside the function for Person type struct, the effects will be
visible from outside the function too. Nothing is being returned from this function. Look at line 13 and 14,
we are converting both the strings to uppercase. Now, let’s see the main function, where we use three different
methods to call upPerson function.

Let’s see the first case. At line 19, we declare a Person type variable pers1 as a value type. In the next two
lines, we fill the fields firstName as Chris and lastName as Woodward of pers1 using . selector. At line 22, we
call the upPerson function on pers1. The upPerson accepts pointer to the Person type, but pers1 is of value type,
so we pass pers1 as: upPerson(&pers1). Because &pers1, is the pointer (address) to pers1.

Let’s see the second case. At line 26, we declare a Person type variable pers2 as a pointer type. In the next two lines,
we fill the fields firstName as Chris and lastName as Woodward of pers1 using the . selector. At line 30, we call the upPerson
function on pers2. The upPerson accepts a pointer to the Person type and pers2 is already of pointer type, so we pass pers2 as
it is: upPerson(pers2).

Let’s see the third case. At line 34, we declare a Person type variable pers3 as a struct literal type as:
pers3 := &Person{"Chris","Woodward"}. At line 35, we call the upPerson function on pers3. As upPerson accepts
a pointer to the Person type and pers3 is already of pointer type, so we pass pers3 as it is: upPerson(pers3).

At lines 23, 31 and 36, we are printing the fields of pers1, pers2, and pers3, respectively. For all three of them,
firstName will now be CHRIS instead of Chris, and lastName will now be WOODWARD instead of Woodward.

*/
