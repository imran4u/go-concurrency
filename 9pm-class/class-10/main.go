package main

import "fmt"

//Call back fuction

func main() {

	fun := func(p *person) *person {
		p.name = "modified"
		p.age = 4
		return p

	}

	personBuilder(fun)

}

type person struct {
	name string
	age  int
}

func personBuilder(fn func(*person) *person) {

	p := &person{name: "imra", age: 20}
	fmt.Println(*p)

	p = fn(p)
	fmt.Println("After modification: ", *p)

}
