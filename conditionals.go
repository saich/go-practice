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

	print("done\n")
}
