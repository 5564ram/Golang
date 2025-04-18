package main

import (
	"fmt"
	"test/min"
)

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := min.IntArray(data) //conversion to type IntArray
	m := min.Min(a)
	fmt.Printf("The minimum of the array is: %v\n", m)
}

func strings() {
	data := []string{"ddd", "eee", "bbb", "ccc", "aaa"}
	a := min.StringArray(data)
	m := min.Min(a)
	fmt.Printf("The minimum of the array is: %v\n", m)
}

func main() {
	ints()
	strings()
}

/*
In min.go in folder min, we implement the Min function (from line 8 to line 18). At line 9, we take min
to be the first element of the collection. Then in the for loop, we traverse the collection comparing
the element at i index with the element at index min_idx. When data.Less(i, min_idx) is true, this means
the ith element is smaller, and we set min to the ith element and min_idx to i as it contains minimum value
encountered so far. After the for loop, min is the smallest item in the collection and we return it.

ElemIx(ix int) interface{} returns the element at the ix index. The function Min and the method ElemIx
return an empty interface{}. This means that they can return an element of any type: we don’t know
beforehand for which collections we want to use the Miner interface.

At line 20, we define an alias IntArray for []int, and at line 25, we define an alias StringArray
for []string.

From line 21 to line 23, we implement Miner for an array of ints. From line 26 to line 28, we do
the same for an array of strings. Now all the work is done!

In main.go we simply test our work. In function ints(), we make an int array data at line 8.
We convert it to a variable a of type IntArray at line 9. Then we can call the Min function
from package min on a, returning the smallest element m from the int array. From line 14 to
line 19, the same steps are repeated for an array of strings.
*/
