package main

import (
	"fmt"
	"sort"
)

func upperBound(arr []int, x int) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i] > x })
}

func lowerBound(arr []int, x int) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i] >= x })
}

func main() {
	arr := []int{1, 3, 3, 5, 7, 9} // Sorted slice
	x := 3

	index := upperBound(arr, x)
	fmt.Println("Upper Bound Index:", index) // Output: 3 (first element greater than 3)

	arr = []int{1, 3, 3, 5, 7, 9} // Sorted slice
	x = 3

	index = lowerBound(arr, x)
	fmt.Println("Lower Bound Index:", index) // Output: 1 (first occurrence of 3)
}

/*
Searching and sorting are very common operations, and the standard library provides these in the package sort.
To sort a slice of ints, import the package sort and simply call the function func Ints(a []int) as:

sort.Ints(arr)
where arr is the array or slice to be sorted in ascending order. To test if an array is sorted, use:

func IntsAreSorted(arr []int) bool
This returns true or false whether or not the array arr is sorted. Similarly, for float64 elements,
you use the following function from sort package :

func Float64s(a []float64)
and for strings, there is a function provided by sort package:

func Strings(a []string)
To search an item in an array or slice, the array must first be sorted (the reason is that the search-functions
are implemented with the binary-search algorithm). Then you can use:

func SearchInts(a []int, n int) int
This searches for n in the slice a, and returns its index. Of course, the equivalent functions for float64s and strings exist also:

func SearchFloat64s(a []float64, x float64) int // search for float64
func SearchStrings(a []string, x string) int // search for strings

https://pkg.go.dev/sort
*/
