package main

import "fmt"

func main() {

LOOP:
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue LOOP
		}
		fmt.Println(i)
	}
}

// it wont go till the infinite loop because of the continue statement in which i will be incremented when i = 5
