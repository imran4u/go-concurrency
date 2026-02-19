package main

import "fmt"

func main() {
	var m map[string]string     // only declear with nil map
	m = make(map[string]string) // if you will not put this line next like will have panic
	m["key"] = "value"
	fmt.Println(m)

	m2 := map[string]string{}
	m2["key2"] = "value-2"
	fmt.Println(m2)

	str := "012345"
	sStr := str[5:6]
	fmt.Println(sStr)
}
