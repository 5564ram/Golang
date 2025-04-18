package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(5))
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(10))
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(11))
}

/*
GOMAXPROCS() : func GOMAXPROCS(n int) int

GOMAXPROCS allows us to set the maximum number of CPUs that can be executed simultaneously.
Calling the function will set the number of CPUs to be n but will return the previous value
set for the number of the CPUs. The default value for this function is the number of CPUs available.
If we input n less than 1, the function will return the previous setting.

runtime.GOMAXPROCS(3) sets the number of CPUs as 3 but returns the previous value 2 set for the number of CPUs.
Furthermore, runtime.GOMAXPROCS(0) just returns the previous setting, which is 3 in this case, because the value
of n is less than 1.

You can modify the number of CPUs using the GOMAXPROCS to enhance the performance of your program.
The number of CPUs needed for a problem depends on the problem. If you set it to 1, you’ll eliminate parallelism
from your program as now the goroutines will just have one CPU for execution and they will have to execute
sequentially despite being parallel in nature.
*/
