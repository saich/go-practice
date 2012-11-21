// functions
package main

import (
	"fmt"
	"math"
)

// A simple function
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
// If I remove the type all together, it fails with compile error
func multiply(x, y int) int {
	return x * y
}

// A function can return any number of results
// This function returns two strings
func swap(x, y string) (string, string) {
	return y, x
}

// Even the return parameters can be named in functions, in which case 
// a return statement without arguments returns the current values of the results.
func swap2(x, y string) (a, b string) {
	a = y
	b = x
	return // This is mandatory even though we are not specifying any value
}

// Functions also support full closures (like in Javacript)
// This function returns a function
func adder() func(int) int {
	sum := 0 // This sum is unique to each adder closure
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Printf("%v + %v = %v\n", 3, 4, add(3, 4))
	fmt.Printf("%v * %v = %v\n", 3, 4, multiply(3, 4))
	a, b := swap("hello", "world")
	fmt.Println("After swap: ", a, b)
	a, b = swap2(a, b)
	fmt.Println("Swap again:", a, b)

	// Functions are values too, (like in Javascript)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
