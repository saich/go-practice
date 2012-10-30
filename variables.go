// variables
package main

import (
	"fmt"
)

// varaiables are declared with var. Type is last parameter
var a, b string

var qwerty bool

// If the variable is initialized, then the type declaration is not mandatory
var x, y = 12, "hey"

// Note: Any incorrect assignment (by type) or less variables are initialized, the compiler throws an error 
// For example, var q, w string = "sa" gives assignment count mismatch 2 = 1

// constants are declared with const.. The type is optional, since it will be deduced automatically based on value
// and hence it is *good practice* not to mention any type for constants
// Constants can be character, string, boolean, or numeric values.
const truth = true

// Numeric constants without type declaration are interesting. They take the type needed by its context
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(a, b, qwerty, x, y)

	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type. 
	// (Outside a function, every construct begins with a keyword and the := construct is not available)
	m, n := false, true
	fmt.Println(m, n)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// needInt(Big) will throw an overflow error

}
