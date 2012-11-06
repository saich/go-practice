// types
package main

import (
	"fmt"
	"math/cmplx"
)

/*
Basic Types in GO:
------------------
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// A struct is a collection of fields.
type Point struct {
	x int
	y int

	// x,y int // This shorthand also works!
}

func main() {
	// %T stands for type
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	// Point object
	point := Point{1, 2}
	fmt.Println(point.x, point.y) // struct fields are accessed using dot
	fmt.Println(point)

	// Go has pointers, (but no pointer arithmetic)
	point2 := &point
	point2.x = 22 // No, we do not use -> in Go. Indirection through pointer is transparent
	fmt.Println(point)
	fmt.Println(&point == point2, point == *point2) // & and * are same as in C++ though..

	// structures can be defined with new also... (similar to C++)
	var point3 *Point = new(Point) // or simply, point3 := new(Point)
	point3.x, point3.y = 23, 56
	fmt.Println(point3)
}
