package main

import "fmt"

func main() {
	sum := 0

	// Go has only one looping construct, the for loop.
	// The basic for loop looks as it does in C or Java, except that 
	// the ( ) are gone (they are not even optional) and the { } are required.
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// As in C or Java, you can leave the pre and post statements empty. (making it equivalent to while loop)
	sum = 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	// infinite loop looks like this
	/*
		for {
			fmt.Println("a")
		}
	*/

}
