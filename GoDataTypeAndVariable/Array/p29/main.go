package main

import "fmt"

func main() {
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	a1 := a[0:4] // ref of subarray {5,6,7,8}

	fmt.Println(a1)
}

/*
[5 6 7 8]
Creating a slice of the array ar as ar[5:7] makes a equal to {5,6}.
Here, the length of a is 2 and the capacity of a is 5. In the last line,
we are reslicing a as a=a[0:4]. Now, a equals to {5,6,7,8}
with a length of 4 and a capacity of 5.
*/
