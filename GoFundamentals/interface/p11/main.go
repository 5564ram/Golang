package main

import "fmt"

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	testFunc := func(any interface{}) { // lambda function in combination with empty interface
		switch v := any.(type) {
		case bool: // if v is bool
			fmt.Printf("any %v is a bool type", v)
		case int: // if v is int
			fmt.Printf("any %v is an int type", v)
		case float32: // if v is float32
			fmt.Printf("any %v is a float32 type", v)
		case string: // if v is string
			fmt.Printf("any %v is a string type", v)
		case specialString: // if v is specialString
			fmt.Printf("any %v is a special String!", v)
		default: // none of types satisfied
			fmt.Println("unknown type!")
		}
	}
	testFunc(whatIsThis)
}

func main() {
	TypeSwitch()
}

/*
any hello is a special String!
*/

/*
In the above code, we define a new type specialString on the basis of string type at line 4. Then,
at line 6, we create a variable whatIsThis of type specialString and give it a value hello. Now,
we have a typeSwitch function at line 8. In the next line, we declare a function testFunc.
This function can take an empty interface as its parameter.

At line 10, we are making a switch statement with cases dependent on the type of v passed to testFunc.

-> The first case will be true if v is of bool type.
-> The second case will be true if v is of int type.
-> The third case will be true if v is of float32 type.
-> The fourth case will be true if val is of string type.
-> The fifth case will be true if v is of specialString type.
Now, the type will be printed, accordingly, when one case is found true.

What if no case is found true? In this case, we have a default case that prints the message of an
unknown type. At line 25, we pass whatIsThis to testFunc, which means that the empty interface is
of type specialString here.

Now, look at main. At line 29, we are calling the TypeSwitch function, which will call the testFunc
for whatIsThis. Line 20 will be executed, and any hello is a special String will be printed on the screen.

------------ Copying a data-slice in a slice of interface{} ------------
Suppose you have a slice of data of myType and you want to put them in a slice of empty interface, like in:
var dataSlice []myType = FuncReturnSlice()
var interfaceSlice []interface{} = dataSlice
This doesn’t work; the compiler gives you the error: cannot use dataSlice (type []myType) as
type []interface{} in assignment. The reason is that the memory layout of both variables is not
the same (try to reason this yourself or see this link). The copy must be done explicitly with
a for-range statement, like in:

var dataSlice []myType = FuncReturnSlice()
var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
for ix, d := range dataSlice {
  interfaceSlice[ix] = d
}
*/
