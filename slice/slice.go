package main

import "fmt"

func main(){
	// s := make([]int, 10)
	s := new([]int)

	fmt.Println("inital slice : ", s)
	fmt.Println("inital size : ", len(*s))
	fmt.Println("inital cap : ", cap(*s))

	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Println("after 10 slice : ", s)
	fmt.Println("after 10 size : ", len(s))
	fmt.Println("after 10 cap : ", cap(s))

}
