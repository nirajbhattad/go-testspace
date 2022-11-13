package gobasics

import (
	"fmt"

	"golang.org/x/tour/pic"
)

/*
 References:
 1) https://www.callicoder.com/golang-slices/

 2) https://rmoff.net/2020/06/25/learning-golang-some-rough-notes-s01e02-slices/

 Length and Capacity of a Slice
	A slice consists of three things -

	A pointer (reference) to an underlying array.
	The length of the segment of the array that the slice contains.
	The capacity (the maximum size up to which the segment can grow).

	Example:
	var a = [6]int{10, 20, 30, 40, 50, 60}
	var s = [1:4]

	The length of the slice is the number of elements in the slice (updated or modified where the lower bound starts), which is 3 in the above example.

	The capacity is the number of elements in the underlying array starting from the first element in the slice. It is 5 in the above example.

	A slice’s length can be extended up to its capacity by re-slicing it. Any attempt to extend its length beyond the available capacity will result in a runtime error.
	Ex: s[1:5] is valid in the above example.
*/

func PlayWithSlices() {
	playWithSlicesEx1()
	playWithSlicesEx2()
	pic.Show(printSlicesPicture)

}

// PlayWithSlices
func playWithSlicesEx1() {

	s := []int{2, 3, 5, 7, 11, 13} // Slice is a pointer to the underlying array
	printSlice(s)
	// len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0] // Re-assigning value based on the pointer against the original array
	printSlice(s)
	// len=0 cap=6 []

	// --
	// Extend its length.
	s = s[:4] // Even after extending, it’s still against the original array that we were pointing to.
	printSlice(s)
	// len=4 cap=6 [2 3 5 7]

	// --
	// Drop its first two values. Increasing the lower bound
	s = s[2:]
	printSlice(s) // Increasing the lower bound ([2:]) we’re actually moving the offset of the pointer against the underlying array.
	// len=2 cap=4 [5 7]
}

// playWithSlices
func playWithSlicesEx2() {
	myArray := [6]int{2, 3, 5, 7, 11, 13}

	y := myArray[:] // No Bounds Specified, defaults to original array
	fmt.Printf("y       len %d\tcap %d\tvalue %v\n", len(y), cap(y), y)
	fmt.Printf("myArray len %d\tcap %d\tvalue %v\n\n", len(myArray), cap(myArray), myArray)
	// y       len 6	cap 6	value [2 3 5 7 11 13]
	// myArray len 6	cap 6	value [2 3 5 7 11 13]

	y = y[0:4] // First four entries, pointing to the same underlying myArray
	fmt.Printf("y       len %d\tcap %d\tvalue %v\n", len(y), cap(y), y)
	fmt.Printf("myArray len %d\tcap %d\tvalue %v\n\n", len(myArray), cap(myArray), myArray)
	// y       len 4	cap 6	value [2 3 5 7]
	// myArray len 6	cap 6	value [2 3 5 7 11 13]

	y = y[:0] // Zero entries, Pointing to same underlying array
	fmt.Printf("y       len %d\tcap %d\tvalue %v\n", len(y), cap(y), y)
	fmt.Printf("myArray len %d\tcap %d\tvalue %v\n\n", len(myArray), cap(myArray), myArray)
	// y       len 0	cap 6	value []
	// myArray len 6	cap 6	value [2 3 5 7 11 13]

	y = y[4:6] // Shift the lower bound, updates the slice lenght and capacity, as we have updated the pointer offset.
	fmt.Printf("y       len %d\tcap %d\tvalue %v\n", len(y), cap(y), y)
	fmt.Printf("myArray len %d\tcap %d\tvalue %v\n\n", len(myArray), cap(myArray), myArray)
	// y       len 2	cap 2	value [11 13]
	// myArray len 6	cap 6	value [2 3 5 7 11 13]

	y = y[1:2] // Shifting the lower bound on the previously updated slice
	fmt.Printf("y       len %d\tcap %d\tvalue %v\n", len(y), cap(y), y)
	fmt.Printf("myArray len %d\tcap %d\tvalue %v\n\n", len(myArray), cap(myArray), myArray)
	// y       len 1	cap 1	value [13]
	// myArray len 6	cap 6	value [2 3 5 7 11 13]

	myArray[5] = 100 // As slice is just a pointer to underlying array, if we change the array, it will reflect in slice.
	fmt.Printf("y       len %d\tcap %d\tvalue %v\n", len(y), cap(y), y)
	fmt.Printf("myArray len %d\tcap %d\tvalue %v\n\n", len(myArray), cap(myArray), myArray)
	// y       len 1	cap 1	value [100]
	// myArray len 6	cap 6	value [2 3 5 7 11 100]

	y[0] = 200 // Conversely, changing the slice value reflects in the array too
	fmt.Printf("y       len %d\tcap %d\tvalue %v\n", len(y), cap(y), y)
	fmt.Printf("myArray len %d\tcap %d\tvalue %v\n\n", len(myArray), cap(myArray), myArray)
	// y       len 1	cap 1	value [200]
	// myArray len 6	cap 6	value [2 3 5 7 11 200]
}

// printSlicesPicture
func printSlicesPicture(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)

	// Init
	for i := range p {
		p[i] = make([]uint8, dx)
	}

	// Sample
	for y := range p {
		for x := range p[y] {
			p[y][x] = (uint8(x) * uint8(y))
		}
	}

	// Sample
	for y := range p {
		for x := range p[y] {
			p[y][x] = (uint8(x) * uint8(y))
		}
	}

	return p
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
