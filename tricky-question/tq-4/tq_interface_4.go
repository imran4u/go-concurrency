package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.height * r.width
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.height + r.width)
}
func main() {

	var s Shape

	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}

	fmt.Println("r==s", r == s)

	// Another way
	r1 := Rect{5.0, 4.0}
	var s1 Shape = r1
	area := s1.Area()
	fmt.Println(area)

}
