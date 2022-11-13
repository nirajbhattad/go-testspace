package gobasics

import "fmt"

/*
	Pointers in Go has to do with three main things:
	1) Garbage Collection
		If a developer uses a pointer; it is the indication about which references need to be reclaimed given its scope.
		It uses something called reference counting to figure that out.

	2) Encapsulation (Information Hiding)
		Pointers are particularly good for structs, which represents complex data structures.
		By returning structs rather than the actual value developers can ensure that only the function that
		created the struct can act upon it.

	3) Immutability
		The concept of interface methods is used, that are nothing more than function pointers to a struct.

	4) Indirection
		Pointers are just indirection. For size, think of loading GB of data.
		Without indirection, you'd have to copy all of it every time.
		Pointer lets you say "here's where to find the data" rather than "here's a copy of the data".

	5) Performance and memory management
       Ex:  a large complex object that you want to use as a parameter to a method.
	   If you don't use a pointer then the entire object is copied and the copy is fed into the method.

	6) No pointer arithmetic in Go.

	References:
	 1. https://www.callicoder.com/golang-pointers/
	 2. https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back
*/

func PlayWithPointers() {
	a := 10
	var b, c *int
	b = &a
	c = &a
	fmt.Println(&a, b, c)
	a = 11
	fmt.Println(&a, b, c)
	fmt.Println(a, *b, *c)

	var m map[int]int
	fn(m)
	fmt.Println(m)
	fmt.Println(m == nil)
}

func fn(m map[int]int) {
	m = make(map[int]int)
	m[10] = 10
}
