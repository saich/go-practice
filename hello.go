package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	// Arguments - Obviously the zeroth argument is the process name
	fmt.Println("Arguments to the program:", os.Args)

	// Note: os.Args.length doesn't work!!
	fmt.Println("Arguments length:", len(os.Args))

	// Printf - Print with formatting.. Verbs @ http://golang.org/pkg/fmt/
	// %v is a safe one - Indicates "print the value in default format"
	fmt.Printf("Pi of Math: %v\n", math.Pi)

	// In Go, a name is exported if it begins with a capital letter.
	// See that in Println, Printf, Pi - All starts with capital letter.
}
