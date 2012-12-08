// http://golang.org/doc/articles/json_and_go.html
package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
	priv int // Private variables will not be used in json.Marshal
}

type Message2 struct {
	alpha string
	Beta  string
}

type A struct {
	name string
}

type B struct {
	NAME string
}

func main() {
	m := Message{"Alice", "Hey...", 35434546546567, 23}
	b, _ := json.Marshal(m) // encode the data
	fmt.Println(string(b))

	var mm, mm3 Message

	json.Unmarshal(b, &mm)  // Unmarshal also takes the pointer to an object, which will be filled
	fmt.Printf("%+v\n", mm) // %+v to print keys + values of a struct

	var mm2 Message2
	json.Unmarshal(b, &mm2) // Unmarshal also takes the pointer to an object, which will be filled
	fmt.Printf("%+v\n", mm2)

	// The field matching can also match case-insensitive fields
	aa := `{"name":"Alice","Body":"Hey...","Time":35434546546567}`
	json.Unmarshal([]byte(aa), &mm3) // Unmarshal also takes the pointer to an object, which will be filled
	fmt.Printf("%+v\n", mm3)         // %+v to print keys + values of a struct

	// During decoding, no private variables of the original struct will be overridden..
	var a A
	json.Unmarshal([]byte(`{"name": "test"}`), &a)
	fmt.Printf("%+v\n", a) // a.name will not be "test" here..

	var bb B
	json.Unmarshal([]byte(`{"name": "test"}`), &bb)
	fmt.Printf("%+v\n", bb) // b.NAME will be "test" here..

	// If we do not know the structure already, we can use interface{} or map[string]interface{}
	var i interface{}
	b = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	json.Unmarshal(b, &i)
	fmt.Println(i)

	// JSON mapping:
	// string -> string, numbers -> float64, null -> nil, bool -> bool

	mmm := i.(map[string]interface{})
	for k, v := range mmm {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is a string:", vv) // TODO: Not sure how, but vv here is representing the value
		case float64:
			fmt.Println(k, "is a number:", vv)
		}
	}

	// TODO: Pretty print the JSON with marshal..

	// TODO: Reference Types while UnMarshalling (and pattern that arises from that)

	// TODO: JSON Streams - Encoder and Decoder

}
