package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("key: %v, value: %v / ", k, v) // read random keys
	}
	keys := make([]string, len(barVal)) // storing all keys in separate slice
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys) // sorting the keys slice
	fmt.Println()
	fmt.Println("\nsorted:")
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v / ", k, barVal[k]) // reading key from keys and value from barVal
	}
}

/*
By default, a map is not sorted, not even on the value of its keys. If you want a sorted map,
copy the keys (or values) to a slice, sort the slice (using the sort package)

In the code above, outside main at line 4, we import sort package to perform sorting, and at line 7,
we make a map barVal. The declaration of barVal shows that its keys will be of string type and values
associated with its keys will be of int type. The initialization is done at the same line. For example,
the key alpha gets value 34, bravo get 56, charlie gets 23 and so on.

In main, we have a for-loop at line 15. This loop prints a key and the value associated with each key
in every iteration, until all key-value pairs are not printed. The output of this for loop will be unsorted,
as maps’ keys are read randomly.

Now at line 18, we make a slice of string keys of length equal to barVal's length ( total number of keys in
barVel). Then we make the for loop at line 20 that reads keys from barVel in random order and stores them one
by one in keys ( at line 21). By the end of for loop, all keys will be present in the keys slice. There is a
great chance that the keys from barVal will not be read in sorted (alphabetical) order and will be stored in
the keys slice. So at line 24, we perform in-place sorting on the strings present in keys using:
sort.Strings(keys).

Again we have a for-loop at line 27. This loop prints a key and the value associated with each key in
every iteration until all key-value pairs are not printed. It reads the key from keys and store it in k.
Then read the value from barVal as: barVal[k].The output of this for loop will be sorted, as barVal's keys
are read from keys slice (containing sorted strings).
*/
