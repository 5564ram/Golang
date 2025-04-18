package main

import (
	"fmt"
	"reflect"
)

type T struct {
	a int "This is a tag"
	b int `A raw string tag`
	c int `key1:"value1" key2:"value2"`
}

func main() {
	t := T{}
	fmt.Println(reflect.TypeOf(t).Field(0).Tag)
	if field, ok := reflect.TypeOf(t).FieldByName("b"); ok {
		fmt.Println(field.Tag)
	}
	if field, ok := reflect.TypeOf(t).FieldByName("c"); ok {
		fmt.Println(field.Tag)
		fmt.Println(field.Tag.Get("key1"))
	}
	if field, ok := reflect.TypeOf(t).FieldByName("d"); ok {
		fmt.Println(field.Tag)
	} else {
		fmt.Println("Field not found")
	}
}

/*
This is a tag
A raw string tag
key1:"value1" key2:"value2"
value1
Field not found

In the above code, at line 4, we import a package reflect. At line 7, we are making a struct T type variable with three fields:
a of type int, b with type int, and c with type int. You may have noticed with all fields, are some tags associated, in quotes and back ticks.

Now, look at main. At line 14, we make a T variable t without assigning fields with values. At line 15, we are printing the tag of field 0 of
t as: fmt.Println(reflect.TypeOf(t).Field(0).Tag). In the next line, we are checking whether we have a field with the name b in t, using the
reflect package as: if field, ok := reflect.TypeOf(t).FieldByName("b"); ok . If b exists then ok becomes true, and the field of b is stored
in field. If ok is true, only then line 17 will be executed. In this case, ok will be true because t has the field b. So the tag of field
b will be printed on the screen.

Similarly, at line 19, we are checking whether we have a field with the name c in t, using the reflect package as:
if field, ok := reflect.TypeOf(t).FieldByName("c"); ok . If c exists, then ok becomes true, and the field of c is stored in field.
If ok is true, only then line 24 will be executed; otherwise, line 26 will be executed. In this case, ok will be false because t has
field d. So, the tag of field d will not be printed on the screen (line 24). Line 26 will be executed, and Field not found will be printed.
*/
