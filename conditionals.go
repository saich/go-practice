package main

import "fmt"

func main() {
	x := 1
	// No () around condition, and {} are must
	if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is not zero")
	}

	// Like for, the if statement can start with a short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if.
	if v := 10; x < v {
		fmt.Println("x is less than", v)
	} else {
		fmt.Println("x is not less than", v) // variables declared in if statement are accessible in else block also
	}

	// Even multiple variables can be declared in the if statement.
	// "else if" is also supported, and you can even declare variables in else if blocks
	// The variable declared in else if block can be accessible in the else block also!
	if min, max := -3, 33; x < min {
		fmt.Println("x is less than", min)
	} else if test := 12; x > max {
		fmt.Println("x is less than ", max)
		// fmt.Println("Test =", test)
	} else {
		fmt.Println("Test2 =", test)
		fmt.Println("x is none")
	}
}
