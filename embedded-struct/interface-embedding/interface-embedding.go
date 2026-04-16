package main

import "fmt"

type Reader interface {
	Read()
}

type A struct{}

func (A) Read() {
	fmt.Println("Read -A")
}

type B struct {
	A
}

func main() {
	b1 := B{}
	b1.Read()

	//
	var r Reader

	var b B
	r = b // ✅ works
	r.Read()

	var pb *B
	fmt.Println("pb==nil", pb == nil)
	// pb.Read() // (*pb).A.Read() -> nil.A.Read() so nil pointer

	r = pb // ✅ works
	// r.Read() // it will give panic: runtime error: invalid memory address or nil pointer dereference
}
