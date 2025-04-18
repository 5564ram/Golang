package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string `The name of the thing`
	field3 int    "How much there are"
}

var tt TagType

func main() {
	tt = TagType{true, "Barack Obama", 1}
	ttType := reflect.TypeOf(tt)
	for i := 0; i < ttType.NumField(); i++ {
		refTag(i)
	}
}

func refTag(ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)     // getting field at a position ix
	fmt.Printf("%v\n", ixField.Tag) // printing tags
}

/*
In the above code, at line 4, we import a package reflect. At line 7, we are making a struct TagType type variable with three fields:
field1 of type bool, field2 with type string, and field3 with type int. You may have noticed with all fields, are some tags associated
in quotes and backticks.

Before studying main let’s focus on the function refTag. See its header at line 20: func refTag(tt TagType, ix int).
The function is taking a Tagtype variable tt and an integer variable ix.

Using the reflect package, we are finding the type of tt and storing the type in ttType (at line 21). Then, we are using the Field function
from the reflect package (at line 22) to find the field at position ix in ttType and storing it as ixField. At line 23, we are using ixField
to find its tag, and then printing it.

Now, look at main. At line 14, we make a TagType variable tt and assign fields with values: true, Barack Obama, and 1. Now, we wish to print
tags associated with fields of tt. We’ll use the package refect for this purpose. As we have three fields, we made a for loop at line 15 with
iterator i, starting from 0 and ending at 2, which means it will iterate three times. In every iteration, we are calling a function refTag,
which takes tt and iterator i. When i is 0, the tag of the first field will be printed. When i is 1, the tag of the second field will be printed.
When i is 2, the tag of the third field will be printed.
*/
