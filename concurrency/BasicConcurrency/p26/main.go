package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("NumCPU is %d\n", runtime.NumCPU())
}

/*
NumCPU: func NumCPU() int
NumCPU() returns the number of logical CPUs that can be used by the current process.z
As you can see, the number of CPUs used by the above code widget is 2.
The above functions taught in this lesson are used in cases where increasing
the number of CPUs will help solve the problem more efficiently
*/
