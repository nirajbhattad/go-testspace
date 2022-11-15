package gobasics

import "fmt"

/*
Interfaces serve a few purposes, but the most common ones are:
	Forcing code encapsulation
	Allowing for more versatile code

Interfaces provide us with a way of writing our own function so that it doesn’t care what type of object you pass
 in as long as it implements the methods required by the interface.

 The empty interface is exactly the same as that interface, except we aren’t declaring any methods that need to be implemented.
 The code is essentially saying “I need an argument, and I don’t care what methods it implements.”

 References:
 	1) https://www.calhoun.io/how-do-interfaces-work-in-go/
	2) https://rmoff.net/2020/06/30/learning-golang-some-rough-notes-s01e05-interfaces/
	3) https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
	4) https://www.callicoder.com/golang-interfaces/

*/

func PlayWithInterface() {
	a := 10
	PrintIt(a)
}

func PrintIt(a interface{}) {
	// a can have any methods. We dont care.
	fmt.Println(a)
}
