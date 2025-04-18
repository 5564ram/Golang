package main

import (
	"fmt"
)

func main() {
	fmt.Println(f(4, 2))
	fmt.Println(f(7, 0))
}

func f(x, y int) (z int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			z = 0
		}
	}()
	return x / y
}

// output:
