package main

import "fmt"

type flt func(int) bool
type slice_split func([]int) ([]int, []int)

func isOdd(integer int) bool { // check if integer is odd
	return integer%2 == 1
}

func isBiggerThan4(integer int) bool { // check if integer is greater than 4
	return integer > 4
}

func filter_factory(f flt) slice_split { // split the slice on basis of func
	return func(s []int) (yes, no []int) {
		for _, val := range s {
			if f(val) {
				yes = append(yes, val)
			} else {
				no = append(no, val)
			}
		}
		return
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("s = ", s)
	odd_even_function := filter_factory(isOdd)
	odd, even := odd_even_function(s)
	fmt.Println("odd = ", odd)
	fmt.Println("even = ", even)

	//separate those that are bigger than 4 and those that are not.
	bigger, smaller := filter_factory(isBiggerThan4)(s)
	fmt.Println("Bigger than 4: ", bigger)
	fmt.Println("Smaller than or equal to 4: ", smaller)
}

/*
The above program has two basic functions. The first function, isOdd takes an integer as a parameter and
returns bool value (see its header at line 7). If the integer is odd, it will return true, otherwise false.
Similarly, the second function isBiggerThan4 takes an integer as a parameter and returns bool value
(see its header at line 15). If the integer is greater than 4, it will return true, otherwise false.

At line 4, we are aliasing a type. A function that takes a single integer as a parameter and returns
a single boolean value is given a type flt. Similarly, at line 5, we are aliasing a type. A function that takes
a slice of integers as a parameter and returns two slices of integers is given a type slice_split.

Now moving towards a major part of the program that is filter_factory function, see its header at line 22.
It takes a function f of type flt(either isOdd or isBiggerThan4). The function filter_factory returns a lambda
function of type split_slices that takes s as a parameter and returns two integer slices yes and no (see line 23).
From line 25 to line 30, we are evaluating s slice on the basis of f function given to filter_factory.
If the f function returns true for an integer, it becomes part of yes, otherwise no.

Let’s see the main function now. At line 36, we declare a slice of integers named s. Then at line38, we call the
filter_factory function with isOdd as the parameter and store the result in odd_even_function of type slice_split.
Now, the odd_even_function is called with s as a parameter, which returns two slices odd and even that contain odd
and even numbers from s, respectively. Similarly at line 43, we call the filter_factory function with isBiggerThan4
as the parameter. You may have noticed we pass s on the same line as: bigger, smaller: = filter_factory(isBiggerThan4)(s),
rather than making a separate split_slices variable and then passing s to it. The result is stored in the bigger and
smaller slices. Printing odd and even slices at line 40 and line 41, respectively, verifies the result. Similarly,
printing bigger and smaller slices at line 44 and line 45 respectively verifies the result.
*/
