package main

import (
	"fmt"
	// "sort"      // this uses the Go sort package, then replace mysort. with sort. in the code below
	"test/mysort" // this uses our own sort package (a subset of the Go sort package)
)

// sorting of slice of integers
func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := mysort.IntSlice(data) //conversion to type IntSlice
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

// sorting of slice of strings
func strings() {
	data := []string{"Monday", "Friday", "Tuesday", "Wednesday", "Sunday", "Thursday", "", "Saturday"}
	a := mysort.StringSlice(data)
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

// a type which describes a day of the week
type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

// sorting of custom type day
func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	Wednesday := day{3, "WED", "Wednesday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	mysort.Sort(&a)
	if !mysort.IsSorted(&a) {
		panic("fail")
	}
	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}

type Person struct {
	firstName string
	lastName  string
}

type Persons []Person

func (p Persons) Len() int { return len(p) }

func (p Persons) Less(i, j int) bool {
	in := p[i].lastName + " " + p[i].firstName
	jn := p[j].lastName + " " + p[j].firstName
	return in < jn
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func personSort() {
	p1 := Person{"Xavier", "Papadopoulos"}
	p2 := Person{"Chris", "Naegels"}
	p3 := Person{"John", "Doe"}
	arrP := Persons{p1, p2, p3}
	fmt.Printf("Before sorting: %v\n", arrP)
	mysort.Sort(arrP)
	fmt.Printf("After sorting: %v\n", arrP)
}

func main() {
	ints()
	strings()
	days()
	personSort()
}

/*
In the file mysort.go:

We make an interface Interface with three methods Len(), Less() and Swap(), as discussed above.
At line 9, we have a function header of Sort that takes Interface as a parameter, and sorts the
data of a collection using the Len() , Less() and Swap() functions.

Then, at line 19, we have a function header of IsSorted that takes Interface as a parameter and
returns true if the data is sorted. Otherwise, false. At line 30, we define a type IntSlice for []int.

From line 32 to line 36, we implement the methods Len(), Less() and Swap() for slice of ints.
At line 38, we define a type StringSlice for []string.

From line 40 to line 45, we implement the methods Len(), Less() and Swap() for slices of strings.
Next, we make wrappers of two functions Sort() and IsSorted() each for IntSlice and StringSlice
from line 48 to line 54.

Now, look at the main.go file. Before studying main, let’s see three basic functions, ints(),
strings(), and days():

Look at the header of function ints() at line 9. We make a slice of integers data and convert it
to type IntSlice by importing: mysort.IntSlice() at line 11, and saving it in a. Then, we call
the Sort function on a. At line 13, we check whether a is sorted or not by calling IsSorted function
on a. If a is not sorted, control will transfer to line 16, and the sorted slice will be printed.
If, sorted then line 14 will be executed.

Now, look at the header of function strings() at line 20. We make a slice of strings data and convert
it to type StringSlice by importing it as: mysort.StringSlice() at line 22, and saving it in a. Then,
we call the Sort function on a. At line 24, we check whether a is sorted or not by calling the IsSorted
function on a. If a is not sorted, control will transfer to line 27, and the sorted slice will be printed.
If sorted, then line 25 will be executed.

Now, before studying the days() function, look at the definition of day type struct at line 31. It has
three fields: one num and integer field and two strings fields shortName and longName. Then, at line 37,
we make another struct of type dayArray with one field data with a type of pointer to slice of type day.
We implement Len(), Less() and Swap() for dayArray type (from line41 to line 43), just like we did for
IntSlice and StringSlice in mysort.go.

Now, look at the header of function days() at line 46. At line 47, we make a day type variable and assign 0
to num, SUN to shortName and Sunday to longName. Then, we make further six day type variables each for a
day (from line 48 to line 53). Next, we make a slice of these seven day type variables and convert it into
dayArray type and store it in a. At line 56, we check whether a is sorted or not by calling the IsSorted
function on a. If a is not sorted, control will transfer to line 60, and the sorted slice will be printed
using longName. If sorted, then line 58 will be executed.

Now, look at main. We are just calling ints(), strings() and days() function at line 67, line 68 and line 69,
respectively. At the end, sorted arrays are shown in output on the screen.
*/
