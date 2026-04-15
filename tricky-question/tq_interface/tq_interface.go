package main

import "fmt"

type S struct {
	a,b,c string
}



func main (){

	x := interface{}(&S{"a","b","c"})
	y := interface{}(&S{"a","b","c"})

	fmt.Println(x == y)

	p := interface{}(S{"a","b","c"})
	q := interface{}(S{"a","b","c"})

	fmt.Println(p == q)

}
