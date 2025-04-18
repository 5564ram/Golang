package main

import (
	"fmt"
)

type List []int

func (l List) Len() int { return len(l) }

func (l *List) Append(val int) { *l = append(*l, val) }

func main() {
	// A bare value
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)\n", lst, lst.Len()) // [1] (len: 1)
	// A pointer value
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)\n", plst, plst.Len()) // &[2] (len: 1)
}

/*
There can be methods attached to the type, and other methods attach a pointer to the type.
However, this does not matter: if for a type T a method Meth() exists on *T and t is a variable
of type T, then t.Meth() is automatically translated to (&t).Meth(). Pointer and value methods
can both be called on the pointer or non-pointer values. This is illustrated in the following program
*/

/*
In the above program, at line 6, we declare a new type List on the basis of the type of
slice on integers([]int). Look at the header of the method Len() at line 8: func (l List) Len() int.
The part (l List) shows that the method can be called by the object of type List only. The method
returns the length of the object l. Now, look at the header of the method Append() at line 10:
func (l *List) Append(val int). The part (l *List) shows that the method can be called by the pointer
to the object of type List only. The method appends the val value at the end of the list l.

Now, look at main, we make a variable lst of type List using var keyword at line 14. In the next line,
we call Append(1) method on lst. At line 16, we are printing lst and length of lst by calling Len() on lst.
We make a variable plst of type List using new() function at line 18. In the next line, we call Append(2)
method on plst. At line 20, we are printing plst and the length of plst by calling Len() on plst.
*/
