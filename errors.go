package main

import (
	"errors"
	"fmt"
	"time"
)

/* In Go, errors are objects that define the *built-in* error interface
type error interface {
	Error() string
}
*/

// In Go, errors are not thrown, they are returned as one of the return values of the function
// @see http://golang.org/doc/go_faq.html#exceptions

// The fmt package knows to call this Error() method when it encounters an error object while printing

// A Custom Error
type MyError struct {
	When time.Time
	What string
}

// Implement the Error method
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// A test function
func run() error {
	return &MyError{time.Now(), "Test Error"} // Return *MyError
}

// A test function - Just return type of the function changed
func run2() *MyError {
	return &MyError{time.Now(), "Test Error 2"} // Return *MyError
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	if err := run2(); err != nil {
		fmt.Println(err)
	}

	//If we are not bothered about custom error types,creating an error is very easy with "errors" package
	err := errors.New("This is an error")

	// To create descriptive error messages, we can use fmt.Errorf
	// %q is like %s with double quotes (and string escape)
	err = fmt.Errorf("user %q (id %q) not found in the system", "sai<>!@$%^&*()_{:}+", "id")

	if err != nil {
		fmt.Println(err)
	}
}
