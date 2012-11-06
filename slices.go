// slices
package main

import (
	"fmt"
)

// Slices are analogous to arrays in other languages, but have some unusual properties.
// A slice points to an array of values and also includes a length.
// []T is a slice with elements of type T.
func main() {
	p := []int{1, 2, 3, 4, 5} // initialize the slice.. The [] comes before the type. Read as slice ([]) of int (type)

	// len gives length of array, cap gives capacity of the slice
	fmt.Println(p, len(p), cap(p))

	// Iterate over the values - same as c++
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}

	// Slices can be re-sliced, creating a new slice value that **points to the same array**.
	// s[lo:hi] evaluates to [lo,hi),i.e; index hi is not included, and contains hi-lo elements 
	// so s[lo:lo] is empty
	fmt.Println("p[1:4] ==", p[1:4])

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])
}
