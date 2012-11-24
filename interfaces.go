// interfaces
package main

import (
	"fmt"
	"math"
	"os"
)

// An interface type is defined by a set of methods. 
// A value of interface can hold any value that implements these methods

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2) // A variable of type MyFloat
	a = f                     // This works since MyFloat implements Abs() function (even though MyFloat never explicitly says that it implements this interface)
	fmt.Println(a.Abs())
	v := Vertex{3, 4}
	a = &v // This works because *Vertex also implements Abs()
	// a = v // This doesn't work since Abs() is not implemented by Vertex (i.e; the method reciever is *Vertex not Vertex)
	fmt.Println(a.Abs())

	var w Writer
	w = os.Stdout

	fmt.Fprintf(w, "Hello Writer!\n")
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// The Reader and Writer interfaces are declared here for demonstration purposes only..
// They are already avaiable in io package
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

// Compound interfaces - An interface can be member of another interface.
type ReadWriter interface {
	Reader
	Writer
}
