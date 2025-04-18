package main

import "fmt"

func main() {
	number := 0

	go func() {
		number++ //reading and modifying the value of 'number'
	}()

	fmt.Println(number) //reading the value of 'number'

}

/*A data race happens when processes have to access the same variable concur­rently i.e.
one process reads from a memory location while another simultaneously writes to the exact same memory location.

We increment the value of the variable number i.e. we first access the value, add 1 to it and then write the new value back to the memory.
number++ takes place using an anonymous go routine. In the next step, we read the value of number and print it onto the console.
However, the output of the code above turns out to be 0 because the main routine finishes itself before the goroutine has a chance to execute itself completely.
We’ll explore more about this concept in the second chapter.

The point to note in the above example is that number++ and fmt.Println(number) are participating in a data race as number++ is reading from and writing
to the same memory location that fmt.Println(number) is reading from.*/
