package main

import (
	"fmt"
	"reflect"
)

type Reader interface {
	Read()
}

type MyReader struct{}

func (m *MyReader) Read() {
	fmt.Println("My reader read")
}

func main() {
	var r Reader
	fmt.Println("r=", r)
	fmt.Println("r==nil", r == nil)
	fmt.Println("Type:", reflect.TypeOf(r))   // Get type using reflect
	fmt.Println("Value:", reflect.ValueOf(r)) // Get value using reflect

	var m MyReader
	r = &m //But whether you use m or &m depends on how the methods of MyReader are defined. notice pointer here : func (m *MyReader) Read() { ...}
	fmt.Println("After .....")
	fmt.Println("r=", r)
	fmt.Println("r==nil", r == nil)
	v, ok := r.(*MyReader)
	fmt.Println("r.Type = ", v, ok)

	//---- Get type and value of interface
	fmt.Println("Type:", reflect.TypeOf(r))   // Get type using reflect
	fmt.Println("Value:", reflect.ValueOf(r)) // Get value using reflect

	fmt.Printf("Type (%%T): %T\n", r)  // Get type using fmt , use capital T not small t.
	fmt.Printf("Value (%%v): %v\n", r) // Get type using fmt

}
