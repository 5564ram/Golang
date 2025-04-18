package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func main() {
	fmt.Println(f())
}

// 2
// The deferred function is executed after the return statement has been evaluated.
// In the function f, we have a deferred function that increments the return value by 1.
// The return value is 1, but the deferred function increments it by 1, so the final return value is 2.
