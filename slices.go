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

	var a []int
	fmt.Println(a == nil, len(a), cap(a)) // The zero value of a slice is nil. The len and cap functions will both return 0 for a nil slice.

	// slices can be initialized by make() function
	// make allocates an array and returns a slice that refers to that array

	a = make([]int, 10, 20) // make takes type, length, capacity..
	fmt.Println(a, len(a), cap(a))

	// Iterate over the values - same as c++
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}

	// Iterate over the indices, values - for.. in loop
	for index, value := range p {
		fmt.Printf("p[%d] = %v, ", index, value)
	}
	fmt.Println()

	// Iterate - If index is not required, we can use _
	// If only index is required, drop ",value" entirely.
	for _, value := range p {
		fmt.Println(value)
	}

	// Slices can be re-sliced, creating a new slice value that **points to the same array**.
	// s[lo:hi] evaluates to [lo,hi),i.e; index hi is not included, and contains hi-lo elements 
	// so s[lo:lo] is empty
	fmt.Println("p[1:4] ==", p[1:4])

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])

	// Both the start and end can be skipped
	fmt.Println(p[:])
	// fmt.Println(p == p[:]) // error - says, slices can only be compared to nil

	s := p[0:]
	s[0] = 99
	fmt.Println(p[0]) // slices point to the same underlying array

	// slices can be created using same syntax from arrays
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	sss := x[:] // a slice referencing the storage of x
	fmt.Println(sss, x)

	sss = x[0:1]
	fmt.Println(sss, len(sss), cap(sss))
	// slices can be grown to upto its capacity, not beyond
	sss = sss[:cap(sss)]
	fmt.Println(sss, len(sss), cap(sss))

	// A slice cannot be grown beyond its capacity. Attempting to do so will cause a runtime panic, 
	// just as when indexing outside the bounds of a slice or array.
	// sss = sss[:20]

	// Similarly, slices cannot be re-sliced below zero to access earlier elements in the array.
	// sss = sss[-1:] // error

	// Growing slices...
	// Create a new slice with make() with larger capacity and copy the values
	qw := make([]int, len(p)*2, (cap(p)+1)*2)
	copy(qw, p) // copy(dest, src)...
	fmt.Println(qw, len(qw), cap(qw))

	// A utility function append(slice, ...) is also avaiable to this job of creating a new slice with larger capacity,
	// if exceeded the current capacity, like c++'s vector
	qw = append(qw, 2, 3)
	fmt.Println(len(qw), cap(qw))
	qw = append(qw, 5656, 64, 45)
	fmt.Println(len(qw), cap(qw))

	// Joining two slices
	temp := append(qw, p...) // The ... denotes to expand the arguments
	fmt.Println(temp)

	// Since the zero value of a slice (nil) acts like a zero-length slice,
	// you can declare a slice variable and then append to it in a loop

	var temp2 []int
	for i := 0; i < 10; i++ {
		temp2 = append(temp2, i)
	}
	fmt.Println(temp2)

	// GOTCHA: Since slices doesn't make copy of underlying array, the complete array
	// will be kept in memory, even though only small part is required.
	// To fix this problem one can copy the interesting data to a new slice before returning it

}
