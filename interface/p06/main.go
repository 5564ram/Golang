package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	val int
}

func (p *Simple) Get() int {
	return p.val
}

func (p *Simple) Set(u int) {
	p.val = u

}

type RSimple struct {
	v int
}

func (p *RSimple) Get() int {
	return p.v
}

func (p *RSimple) Set(u int) {
	p.v = u
}

func fI(it Simpler) int {
	switch it.(type) {
	case (*Simple):
		it.Set(5)
		return it.Get()
	case (*RSimple):
		it.Set(10)
		return it.Get()
	default:
		return 99

	}
}

func main() {
	var s Simple
	fmt.Println(fI(&s)) // &s is required because Get() is defined with the type pointer as receiver
	var r RSimple
	fmt.Println(fI(&r))
}
