package main

import (
	"fmt"
	"math"
)

// GO do not have classes. But, we can define methods (member functions) on structs / custom types
type Vertex struct {
	x, y float64
}

// This is a method declared on Vertex type..
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
	// Note that there is no this... In place of 'this', we control the variable name
	// that refers to our object in the function declaration

	// In this example, v is the "method receiver"
}

func (v *Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// If the member function begins with capital letter, consider it public, else private method

// Methods can also be defined on non-structs (basically any type declared in this package)
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f) // Go doesn't support automatic type casting.. 
	}
	return float64(f)

	// I'd prefer to write this function using a ternary operator, but Go doesn't have ternary operator
}

// Method receivers can be named type or pointer to named type.
// If type is used, a copy of the object is made with each method call. Any changes made
// to this function will have NO EFFECT on the original value

// If pointer to type is used, no copy will be made, and changes happen to the original object itself

func (v Vertex) Modify1() {
	v.x *= 10
	v.y *= 5
}

func (v *Vertex) Modify2() {
	v.x *= 10
	v.y *= 5
}

func main() {
	v := Vertex{2, 3}
	fmt.Println(v.Abs()) // Public method
	fmt.Println(v.abs()) // Private method

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// call to Modify1 will not have any effect on v's value
	// since the changes will be made on a copy which is passed to the method Modify1
	fmt.Println("v before Modify1", v)
	v.Modify1()
	fmt.Println("v after Modify1", v)

	fmt.Println("v before Modify2", v)
	v.Modify2()
	fmt.Println("v after Modify2", v)

}
