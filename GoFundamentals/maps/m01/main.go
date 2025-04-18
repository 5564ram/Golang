package main

import "fmt"

func main() {
	mapLit := map[string]int{"one": 1, "two": 2} // making map & adding key-value pair
	var mapAssigned map[string]int
	mapCreated := make(map[string]float32) // making map with make()
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5 // creating key-value pair for map
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3 // changing value of already existing key
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}

/*

Maps are cheap to pass to a function because only a reference is passed (so 4 bytes on a 32-bit machine,
8 bytes on a 64-bit machine, no matter how much data they hold)

In the code above, in main at line 5, we make a map mapLit and add key value pairs to it.
The declaration of mapLit shows that its keys will be of string type and the values associated
with its keys will be of int type. You can see we make two keys named one and two. The values
associated with one is 1 and with two is 2.

Similarly, at line 6, we make a map mapAssigned. The declaration of mapAssigned shows that its
keys will also be of string type and the values associated with its keys will be of int type too.
In the next line, we assign values to mapLit.

At line 7, we create a map mapCreated using the make function, which is equivalent to
mapCreated := map[string]float{}. The declaration of mapCreated shows that its keys will be of string
type and the values associated with its keys will be of float32 type. Then in the next line,
we assign mapAssigned with the map of mapLit, which means that mapLit and mapAssigned now possess
the same pairs.

At line 9 and line 10, we are making key-value pairs (each pair line by line) for mapCreated.
At line 9 we create key key1 and give the value 4.5 to it. At line 10 we create the key key2 and
give the value 3.14159 to it. Now at line 11, we are changing the value associated with the key of
mapAssigned, already present before. We are giving the key two a new value of 3.

Now, from line 12 to line 15, we are printing certain values to verify the maps’ behaviors.
At line 12, we are printing the value associated with key one of mapLit which is 1. In the next
line, we are printing the value associated with key key2 of mapCreated, which is 3.14159. At line 14,
we are printing the value associated with key two of mapLit, which is 3 (because of reference, as the
value changes at line 11). In the last line, we are trying to print the value associated with the key
ten of mapLit, which doesn’t even exist. So 0 will be printed because mapLit contains int value types,
and the default value for an integer is 0.

Note: Don’t use new, always use make with maps.

Unlike arrays, maps grow dynamically to accommodate new key-values that are added; they have no
fixed or maximum size. However, you can optionally indicate an initial capacity cap for the map, as in:

make(map[keytype]valuetype, cap)
For example:
map2 := make(map[string]float, 100)
This will create a map with an initial capacity of 100 key-value pairs.
When the map has grown to its capacity, and a new key-value is added, then the size of the map
will automatically increase by 1. So for large maps or maps that grow a lot, it is better for
performance to specify an initial capacity; even if this is only known approximately.
If you don’t specify the initial capacity, a default capacity of 8 bytes is taken (on a 64-bit machine).
*/
