package main

import "fmt"

func main() {

	for i := range 6 {
		if i == 3 {
			goto label
		}
		fmt.Println(i)
	}
label:
	fmt.Println("I am in label")

}
