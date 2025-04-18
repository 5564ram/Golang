package main

import "fmt"

type Any interface{}
type Anything struct{}

func main() {
	any := getAny()
	if any == nil {
		fmt.Println("any is nil")
	} else {
		fmt.Println("any is not nil")
	}
}

func getAny() Any {
	return getAnything()
}
func getAnything() *Anything {
	return nil
}

/*
When getAnything() returns nil, it is assigned to the Any interface in getAny().
In Go, an interface value consists of two parts:
A type (the dynamic type of the value it holds).
A value (the actual value).
When getAnything() returns nil, the Any interface still has a type (*Anything) and a value (nil).
As a result, the Any interface is not nil because it has a type (*Anything), even though the value is nil.
Output:

The if any == nil check fails because the Any interface is not truly nil (it has a type).
Therefore, the program prints:
any is not nil
*/
