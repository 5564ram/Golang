package main

import (
	"fmt"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43,
		"lima": 98}
)

func main() {
	invMap := make(map[int]string, len(barVal)) // interchanging types of keys and values
	for k, v := range barVal {
		invMap[v] = k // key becomes value and value becomes key
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("key: %v, value: %v / ", k, v)
	}
}

/*
In the above code, outside main at line 6, we make a map barVal. The declaration of barVal shows that
its keys will be of string type and values associated with its keys will be of int type. The initialization
is done at the same line. For example, the key alpha gets value 34, bravo get 56, charlie gets 23 and so on.

In main, at line 13, we make a map invMap. The declaration of invMap shows that its keys will be of int
type and values associated with its keys will be of string type. The length of invMap is the same as barVal,
and the map invMap is the inverse of barVal. That is why the type of keys and values are interchanged between
them.

Then we have the for loop at line 14, where we read barVal key in k and value for each k in v for each iteration.
To invert a map, the key becomes value and the value becomes key. This means that invMap[v]=k( see line 15)
will serve the purpose. As v was the value associated to k key of barVal, to invert the map k should be the
value associated with v key of invMap. Again we have a for-loop at line 18. This loop prints a key and the
value associated with each key in every iteration until all key-value pairs are not printed.

Because keys must be unique, the code goes wrong when the original value items are not unique.
In that case, no error occurs, but the inverted map will not contain all pairs from the original map.
One solution is to carefully test for uniqueness and using a multi-valued map. In this case, use a map
of type map[int][]string.
*/
