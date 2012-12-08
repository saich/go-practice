// generic-interface
package main

import (
	"fmt"
)

// The interface{} (empty interface) type describes an interface with zero methods.
// Every Go type implements at least zero methods and therefore satisfies the empty interface.

// The empty interface serves as a general container type

type Any interface{}

func main() {
	var any Any
	fmt.Printf("%T, %v\n", any, any) // %T prints the type
	any = 1
	fmt.Printf("%T, %v\n", any, any)
	any = 3.23434
	fmt.Printf("%T, %v\n", any, any)
	any = "A string"
	fmt.Printf("%T, %v\n", any, any)

	// A type assertion accesses the underlying concrete type. Consider this type assertion as C++'s dynamic_cast
	var str string = any.(string)
	fmt.Printf("%v\n", str)

	// The type can also be accessed, but it is allowed only in switch statement
	switch any.(type) {
	case int:
		fmt.Println("any is int")
	case string:
		fmt.Println("any is string")
	}

	// If the generic type doesn't match the underlying concrete type, it panics
	var i int = any.(int)
	fmt.Printf("%v\n", i)
}
