package main

import (
	"fmt"
	"oddEvenTest/even"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the integer %d even : %v\n", i, even.Even(i))
	}
}
