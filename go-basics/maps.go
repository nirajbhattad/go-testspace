package gobasics

import (
	"fmt"
	"unsafe"
)

/*
Maps are reference types. When you assign a map to a new variable, they both refer to the same underlying data structure.
The same concept applies when you pass a map to a function. Any changes done to the map inside the function is also visible to the caller.

A map value is a pointer to a runtime.hmap structure.

When you write the statement m := make(map[int]int)

# The compiler replaces it with a call to runtime.makemap, which has the signature

makemap implements a Go map creation make(map[k]v, hint)

If the compiler has determined that the map or the first bucket
can be created on the stack, h and/or bucket may be non-nil.
If h != nil, the map can be created directly in h.
If bucket != nil, bucket can be used as the first bucket.

func makemap(t *maptype, hint int64, h *hmap, bucket unsafe.Pointer) *hmap

# The type of the value returned from runtime.makemap is a pointer

Maps, like channels, but unlike slices, are just pointers to runtime types.
As you saw above, a map is just a pointer to a runtime.hmap structure.

Go maps are not goroutine safe, you must use a sync.Mutex, sync.RWMutex or other memory barrier primitive
to ensure reads and writes are properly synchronised.
Getting your locking wrong will corrupt the internal structure of the map.

Of all of Go’s built in data structures, maps are the only ones that move data internally. When you insert or delete entries,
the map may need to rebalance itself to retain its O(1) guarantee. This is why map values are not addressable.

Maps are often used for shared state
Maps are more complex structures
Maps move things
Go’s map is a hashmap - Hash function needs to be stable, distribution and collision resistant(avoiding poor distribution).
*/
func PlayWithMaps() {

	var m map[int]int
	var p uintptr
	fmt.Println(unsafe.Sizeof(m), unsafe.Sizeof(p)) // 8 8 (linux/amd64)

	var m1 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	var m2 = m1
	fmt.Println("m1 = ", m1)
	fmt.Println("m2 = ", m2)

	m2["ten"] = 10
	fmt.Println("\nm1 = ", m1)
	fmt.Println("m2 = ", m2)

	// 	# Output
	// m1 =  map[one:1 two:2 three:3 four:4 five:5]
	// m2 =  map[one:1 two:2 three:3 four:4 five:5]

	// m1 =  map[one:1 two:2 three:3 four:4 five:5 ten:10]
	// m2 =  map[one:1 two:2 three:3 four:4 five:5 ten:10]
}
