package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x)) // float64
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)                      //3.4
	fmt.Println("type:", v.Type())                //float64
	fmt.Println("kind:", v.Kind())                //float64
	fmt.Println("value:", v.Float())              // 3.4
	fmt.Println(v.Interface())                    // 3.4
	fmt.Printf("value is %5.2e\n", v.Interface()) //3.40
	y := v.Interface().(float64)
	fmt.Println(y)
}
