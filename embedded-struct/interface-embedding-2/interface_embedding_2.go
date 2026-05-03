package main

import "fmt"

type Read interface {
	ReadF()
	Common()
}

type Write interface {
	WriteF()
	Common() 
}

type ReadWrite interface {
	Read
	Write
}

type ReadWriteImpl struct{}

func (rw *ReadWriteImpl) ReadF() {
	fmt.Println("Read-write ReadF")
}

func (rw *ReadWriteImpl) WriteF() {
	fmt.Println("Read-write WriteF")
}

//If multiple embedded interfaces have the same method signature, Go deduplicates them automatically.
func (rw *ReadWriteImpl) Common() {
	fmt.Println("Read-write Common")
}

/*
⚠️ When would it FAIL?

If method signatures differ:

type A interface {
	Common() string
}

type B interface {
	Common() int
}

type C interface {
	A
	B
}

👉 ❌ Compile-time error: duplicate method Common
*/



func main() {
	var rw ReadWrite
	rw = &ReadWriteImpl{}
	rw.ReadF()
	rw.WriteF()

	var r Read
	r = &ReadWriteImpl{}
	r.ReadF()
	r.Common()

	var w Write
	w = &ReadWriteImpl{}
	w.WriteF()
	r.Common()

}
