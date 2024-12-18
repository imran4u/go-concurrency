package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)
	fmt.Println("len : ", len(slice), "Cap : ", cap(slice))

	slice1 := new([]int)
	fmt.Println("len : ", len(*slice1), "Cap : ", cap(*slice1))

}
