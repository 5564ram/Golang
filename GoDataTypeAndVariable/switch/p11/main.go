package main

import "fmt"

func main() {
	var num1 int = 100

	// switch statement with values
	switch num1 {
	case 98, 99: // first case: num1 = 98 or 99
		fmt.Println("It's equal to 98")
	case 100: // second case: num1 = 100
		fmt.Println("It's equal to 100")
	default: // optional/ default case
		fmt.Println("It's not equal to 98 or 100")
	}

	// switch statement with conditions
	num1 = 100
	switch {
	case num1 < 0:
		fmt.Println("Number is negative")

	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")

	default:
		fmt.Println("Number is 10 or greater")
	}

	// Initialization within the switch statement
	switch num1 := 100; {
	case num1 < 0:
		fmt.Println("Number is negative")

	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")

	default:
		fmt.Println("Number is 10 or greater")
	}
}
