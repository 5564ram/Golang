package main

func main() {
	// const i = 5
	// ptr1 := &i  //error: cannot take the address of i
	// ptr2 := &10 //error: cannot take the address of 10
	var p *int = nil // Making a nil pointer
	*p = 0
	// panic: runtime error: invalid memory address or nil pointer dereference
}

/*
Oops! The program crashed. The reason is that we create a nil pointer at line 4 and then try to dereference it at line 5.
This is not allowed. Remember that a nil pointer dereference is illegal and makes a program crash.
*/
