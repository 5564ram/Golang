package main

import "fmt"

const (
	A = iota
	B
	C
	D
)
const (
	E = iota
	F
	G
	H
)

func main() {
	fmt.Println(A, B, C, D)
	fmt.Println(E, F, G, H)
}
