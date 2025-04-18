package main

import "fmt"

type T1 struct {
	a int
}

type T2 struct {
	b int
}

type T3 struct {
	a int
}

func main() {
	t1 := T1{10}
	t2 := T2{10}
	t3 := T3{10}
	t4 := T3{20}
	fmt.Println(t1, t2, t3, t4)
	// fmt.Println("t1 == t2?", t1==t2) // <-- invalid operation: t1 == t2 (mismatched types T1 and T2)
	// fmt.Println("t1 == t3?", t1==t3) // <-- invalid operation: t1 == t3 (mismatched types T1 and T3)
	fmt.Println("t1 == t3?", t1 == T1(t3)) // true
	fmt.Println("t3 == t4?", t3 == t4)     // false
}

/*
In the above code, at line 4, we make a struct of type T1 with one int field a. Then at line 8, we make a
struct of type T2 with one int field b. Again at line 12, we make a struct of type T3 with one int field a.

Now, look at main. At line 17, we make a T1 type variable t1 using struct_literal assigning a value of 10.
In the next line, we make a T2 type variable t2 using struct_literal assigning b value of 10. Then, at line 19,
we make a T3 type variable t3 using struct_literal, assigning a a value of 10. Similarly, in the next line,
we make a T3 type variable t4 using struct_literal, assigning a a value of 20. At line 21, we are printing
t1, t2, t3 and t4 to verify the outputs.

Now, let’s perform some comparisons between these structs. See the commented line 22. At line 22, comparison between
t1 and t2 was made via the == operator. It will give an error because the fields of t1 and t2 are not the same.
Similarly, see the next commented line, where the comparison between t1 and t3 was made. It will give an error not
because fields of t1 and t3 are not the same. They are the same. However, you have to convert the type first (see line 24).
The struct t3 is first type-casted to T1 and then compared with t1 as: t1==T1(t3), which will give true because both
fields (a) have a value of 10. In the last line, we are comparing t3 with t4. This doesn’t even type conversion because
t3 and t4 are of type T3. This will give false because a of t3 is 10 and a of t4 is 20.
*/
