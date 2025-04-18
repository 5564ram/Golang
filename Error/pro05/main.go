package main

import "fmt"

func main() {
	f(5)
}
func f(n int) {
	defer func() { fmt.Println(n) }()
	if n == 0 {
		panic(0)
	}
	f(n - 1)
}

// Output:
