package main

import (
	"fmt"
	"unsafe" // to use function Sizeof()
)

type T1 struct {
	a, b int64
}

type T2 struct {
	t1 *T1 // pointer to T1
}

type T3 struct {
	t1 T1 // value type of T1
}

func main() {
	fmt.Println("Size of T1:", unsafe.Sizeof(T1{}))      // T1 value type
	fmt.Println("Size of T2:", unsafe.Sizeof(T2{&T1{}})) // T2 containing pointer to T1
	fmt.Println("Size of T3:", unsafe.Sizeof(T3{}))      // Value of T3
}

/*
Size of T1: 16
Size of T2: 8
Size of T3: 16

If you have a struct type T and you quickly want to see how many bytes a value occupies in memory, use:

size := unsafe.Sizeof(T{})
There is a difference between the size of a struct containing another struct and the size of a struct containing a pointer to another struct.

In the program above, at line 4, we import the package unsafe to use the function Sizeof() that tells the size of the parameter passed to it.
At line 7, we declare a struct of type T1, which has two fields of type int64; a and b. At line 11, we declare another struct of type T2,
which has one field, and is a pointer to the variable of type T1 called t1. At line 15, we declare another struct of type T3 with one field t1,
which is a value of type T1.
Now, look at the main. At line 20, we are printing the size of T1. The size of one int64 type number is 8 bytes. Since the struct T1 has two
int64 type numbers, the total size will be 16.
At line 21, we are printing the size of T2{&T1{}}. Since it’s a pointer or address type, the size of the pointer is the same, regardless of
what data type they are pointing to. On a 32-bit machine, the size of the pointer is 4 bytes, where on a 64-bit machine, the size of the
pointer is 8 bytes. In this case, the size of the pointer is also 8 bytes.
At line 22, we are printing the size of T3. Where T3 holds a value of type T1, and the size of T1 is 16. So, the size of T3 will also be
16 bytes.
*/
