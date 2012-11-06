// maps
package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

// map maps key => value (string => Vertex)
var m map[string]Vertex

var n map[int]Vertex = make(map[int]Vertex) // Also, initialize

var o map[Vertex]string = make(map[Vertex]string)

// Initialize with data
var mp = map[string]Vertex{
	"New Delhi": Vertex{0, 0},
	"Hyderabad": Vertex{10, 10}, // NOTE: This comma here is mandatory, to put the closing } in next line
}

// The typename can be omitted also..
var mp2 = map[string]Vertex{
	"New Delhi": {0, 0},
	"Hyderabad": {10, 10},
}

func main() {
	fmt.Println(m) // maps by default are initialized to nil
	fmt.Println(m == nil)

	m = make(map[string]Vertex) // Initialize a map.. Without this, entries cannot be inserted
	n[32] = Vertex{1, 2}        // Inserting an element
	n[48] = n[32]

	m["Taj Mahal"] = Vertex{23, 45}
	fmt.Println(m)
	fmt.Println(m["Taj Mahal"]) // Retrieve an element

	delete(m, "Taj Mahal")      // If the key doesn't exist, its a no-op
	fmt.Println(m["Taj Mahal"]) // In case the key is not present, a zeroed value will be returned

	// Check if a key exists in the map
	// _ here signifies that we do not need this return value
	_, ok := m["Taj Mahal"]
	fmt.Println("Taj Mahal exists:", ok)
	a := Vertex{11, 22}
	b := Vertex{11, 22}

	o[a] = "sai"
	fmt.Println(o[b]) // surprisingly, structs are compared by value

	// Iterate over all elements in map..
	// NOTE: Unlike C++, order of iteration is undefined.. The order in which elements are retrieved
	// is unknown, and can even change for accesses on same data
	// source: https://groups.google.com/d/msg/golang-nuts/YfDxpkI34hY/4pktJI2ytusJ
	for k, v := range n {
		fmt.Println(k, v)
	}

	// Number of elements in the map
	fmt.Println(len(n), len(m), len(o))
}
