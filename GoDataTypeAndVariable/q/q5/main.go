package main

import "fmt"

func main() {

LOOP:
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto LOOP
		}
		fmt.Println(i)
	}
}

// it will go till the infinite loop because of the goto statement in which i wont be incremented when i = 5
