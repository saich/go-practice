package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// A case body breaks automatically (unlike C++ which requires break statement)
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
	case "linux":
		fmt.Println("Linux")
	}

	// fallthrough statement makes a fallthrough to the next case/default statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
		fallthrough
	case "linux":
		fmt.Println("Linux")
		fallthrough
	default:
		fmt.Println(os)
	}

	// Unlike C++, swtich case conditions can be expressions (instead of just functions)

	fmt.Printf("When's Saturday? ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("Day after tomorrow")
	default:
		fmt.Println("Too far")
	}

	// The case statements shall be evaluated in same order (top-to-bottom), 
	// and stops when a case condition succeeds

	fc := func() int {
		fmt.Println("fc function called")
		return 100
	}
	i := 0
	switch i {
	case 0:
	case fc(): // This won't be called since i == 0
	}

	// switch without condition is equal to switch true {}
	// This can be used in place of long if-elseif-else blocks
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// TODO: Multiple switch conditions can be given in a case statement

}
