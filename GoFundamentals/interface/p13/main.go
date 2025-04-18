package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}

/*
type: float64
value: 3.4
type: float64
kind: float64
value: 3.4
3.4
value is 3.40e+00
3.4
*/
/*
In the above code at line 4, we import a package reflect. Look inside main. At line 8, we declare a
float64 type variable x and assign a 3.4 value to it. Now, in the next line, we are printing the
type of x through the TypeOf() function of the reflect package which is float64. At line 10, we
are reading the value of x in new variable v through ValueOf() function of the reflect package
and printing it in the next line which is 3.4.
Look at line 12. The type of v is printed just by using the Type() method, which is float64.
In the next line, the Kind() method is called by v, and the returned value is printed, which
is also the type of v, i.e., float64. At line 14, to print the value of v, the Float() method
is called by v. At line 15, we recover the interface value of v, and print v in the %5.2e format
in the next line. At line 17, we convert the type of recovered value to float64 again and store
it in y, which is later printed in the next line.
*/
