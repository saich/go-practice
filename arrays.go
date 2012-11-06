// arrays
package main

import (
	"fmt"
)

func main() {

	// An array's size is fixed; its length is part of its type ([4]int and [5]int are distinct, incompatible types). 
	// Arrays can be indexed in the usual way, so the expression s[n] accesses the nth element

	var a [4]int
	a[0] = 1
	i := a[0]
	// i == 1
	fmt.Println(i)

	// Arrays do not need to be initialized explicitly; the zero value of an array is a ready-to-use array whose elements are themselves zeroed:
	fmt.Println(a[2] == 0)
	fmt.Println("Hello World!")

	// An array can be initiazed this way
	b := [2]string{"Hello", "World"}
	fmt.Println(b)

	// c := [2]int{1, 2, 3} // This will give out of bounds error at compile time

	// This will work fine.. d[1] wil be equal to 0
	d := [2]int{45}
	fmt.Println(d)

	// We can also let the compiler count the array size for us..
	e := [...]int{1, 2, 3, 4, 5}

	// For arrays, length and capacity will be equal, since any array wouldn't need more space than its length
	fmt.Println(len(e), cap(e))

	// Unlike arrays in C++, arrays are values (just like int, float etc.)
	// The below line makes a copy of the array, not just a reference
	f := e
	f[0] = 12
	fmt.Println(e, f) // e != fallthrough

	// Arrays are passed by value (i.e; a copy is made) when passing to a function
	// TO avoid the copy and pass by ref, we need to pass the pointer to the array instead of the array directly

	// Since arrays in Go are a bit inflexible, slices are used everywhere. Slices are built on 
	// arrays and provides greater power and convinience.
}
